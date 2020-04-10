package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ontio/saga/config"
)

type SagaDB struct {
	db *gorm.DB
}

var DefDB *SagaDB

func NewDB() (*SagaDB, error) {
	db, dberr := gorm.Open("mysql", config.DefConfig.ProjectDBUser+
		":"+config.DefConfig.ProjectDBPassword+
		"@tcp("+config.DefConfig.ProjectDBUrl+
		")/"+config.DefConfig.ProjectDBName+
		"?charset=utf8")
	if dberr != nil {
		return nil, fmt.Errorf("[NewSagaDB] open db error: %s", dberr)
	}
	return &SagaDB{
		db: db,
	}, nil
}

func (this *SagaDB) QueryRequestNum() (int, error) {
	return 0, nil
}

func (this *SagaDB) SearchApi(key string) {

}

func (this *SagaDB) Close() {
	this.db.Close()
}
