package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/saga/cmd"
	"github.com/ontio/saga/config"
	"github.com/ontio/saga/dao"
	"github.com/ontio/saga/restful"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "Bonus CLI"
	app.Action = startSaga
	app.Version = config.Version
	app.Copyright = "Copyright in 2018 The Ontology Authors"
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.RestPortFlag,
		cmd.NetworkIdFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		cmd.PrintErrorMsg(err.Error())
		os.Exit(1)
	}
}

func startSaga(ctx *cli.Context) {
	initLog(ctx)
	if err := initConfig(ctx); err != nil {
		log.Errorf("[initConfig] error: %s", err)
		return
	}
	if err := initDB(ctx); err != nil {
		log.Errorf("[initDB] error: %s", err)
		return
	}
	restful.StartServer()
	waitToExit()
}

func initLog(ctx *cli.Context) {
	//init log module
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	logName := fmt.Sprintf("%s%s", config.LogPath, string(os.PathSeparator))
	log.InitLog(logLevel, logName, log.Stdout)
}

func initConfig(ctx *cli.Context) error {
	//init config
	return cmd.SetOntologyConfig(ctx)
}

func initDB(ctx *cli.Context) error {
	if config.DefConfig.NetWorkId == config.NETWORK_ID_MAIN_NET {
		userName, err := getDBUserName()
		if err != nil {
			return fmt.Errorf("getDBUserName failed, error: %s", err)
		}
		pwd, err := getDBPassword()
		if err != nil {
			return fmt.Errorf("getDBPassword failed, error: %s", err)
		}
		config.DefConfig.ProjectDBUser = userName
		config.DefConfig.ProjectDBPassword = string(pwd)
	}
	var err error
	dao.DefDB, err = dao.NewDB()
	if err != nil {
		return err
	}
	return nil
}

func getDBUserName() (string, error) {
	fmt.Printf("DB UserName:")
	var userName string
	n, err := fmt.Scanln(&userName)
	if n == 0 {
		return "", fmt.Errorf("db username is wrong")
	}
	if err != nil {
		return "", err
	}
	return userName, nil
}

// GetPassword gets password from user input
func getDBPassword() ([]byte, error) {
	fmt.Printf("DB Password:")
	passwd, err := gopass.GetPasswd()
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {

			log.Infof("bonus server received exit signal: %s.", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}
