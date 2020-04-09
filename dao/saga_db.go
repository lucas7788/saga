package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ontio/saga/config"
	"fmt"
)


func NewDB() (*sql.DB, error) {
	db, dberr := sql.Open("mysql",
		config.DefConfig.ProjectDBUser+
			":"+config.DefConfig.ProjectDBPassword+
			"@tcp("+config.DefConfig.ProjectDBUrl+
			")/"+config.DefConfig.ProjectDBName+
			"?charset=utf8")
	if dberr != nil {
		return nil, fmt.Errorf("[NewSagaDB] open db error: %s",dberr)
	}
	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[NewSagaDB] ping failed: %s",err)
	}
	return db, nil
}
