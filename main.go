package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/ontio/ontology/common/log"
	"runtime"
	"github.com/urfave/cli"
	"github.com/ontio/saga/config"
	"fmt"
	"github.com/ontio/saga/restful"
	"github.com/ontio/saga/cmd"
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
	restful.StartServer()
	waitToExit()
}
func initLog(ctx *cli.Context) {
	//init log module
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	logName := fmt.Sprintf("%s%s", config.LogPath, string(os.PathSeparator))
	log.InitLog(logLevel, logName, log.Stdout)
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