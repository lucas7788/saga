package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ontio/saga/config"
	"github.com/ontio/saga/models"
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
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return &SagaDB{
		db: db,
	}, nil
}

func (this *SagaDB) Init() error {
	if !this.db.HasTable(&models.BuyRecord{}) {
		db := this.db.CreateTable(&models.BuyRecord{})
		if db.Error != nil {
			return db.Error
		}
		this.db = db
	}
	if !this.db.HasTable(&models.APIInfo{}) {
		db := this.db.CreateTable(&models.APIInfo{})
		if db.Error != nil {
			return db.Error
		}
		this.db = db
	}
	if !this.db.HasTable(&models.APIKey{}) {
		db := this.db.CreateTable(&models.APIKey{})
		if db.Error != nil {
			return db.Error
		}
		this.db = db
	}
	return nil
}

func (this *SagaDB) DeleteTable() {
	this.db.DropTableIfExists(&models.BuyRecord{})
	this.db.DropTableIfExists(&models.APIInfo{})
	this.db.DropTableIfExists(&models.APIKey{})
}

func (this *SagaDB) InsertApiInfo(apiInfo *models.APIInfo) error {
	db := this.db.Create(apiInfo)
	if db.Error != nil {
		return db.Error
	}
	this.db = db
	return nil
}

func (this *SagaDB) QueryPriceByApiId(ApiId int) (string, error) {
	info := &models.APIInfo{}
	db := this.db.Table("api_infos").Find(info, "api_id=?", ApiId)
	if db.Error != nil {
		return "", db.Error
	}
	return info.ApiPrice, nil
}

func (this *SagaDB) InsertBuyRecord(buyRecord *models.BuyRecord) error {
	db := this.db.Create(buyRecord)
	if db.Error != nil {
		return db.Error
	}
	this.db = db
	return nil
}

func (this *SagaDB) InsertApiKey(apiKey *models.APIKey) error {
	db := this.db.Create(apiKey)
	if db.Error != nil {
		return db.Error
	}
	this.db = db
	return nil
}

func (this *SagaDB) QueryRequestNum(apiKey string) (int, error) {
	key := &models.APIKey{}
	db := this.db.Table("api_keys").Find(key, "api_key=?", apiKey)
	if db.Error != nil {
		return 0, db.Error
	}
	this.db = db
	return key.UsedNum, nil
}

func (this *SagaDB) QueryApiInfoByPage(start, pageSize int) (infos []models.APIInfo, err error) {
	db := this.db.Table("api_infos").Limit(pageSize).Find(&infos, "id>=?", start)
	if db.Error != nil {
		return nil, db.Error
	}
	return
}

func (this *SagaDB) SearchApi(key string) ([]models.APIInfo, error) {
	var info []models.APIInfo
	k := "%" + key + "%"
	db := this.db.Table("api_infos").Where("api_desc like ?", k).Find(&info)
	if db.Error != nil {
		return nil, db.Error
	}
	return info, nil
}

func (this *SagaDB) VerifyApiKey(apiKey string) error {
	key := &models.APIKey{}
	db := this.db.Table("api_keys").Find(key, "api_key=?", apiKey)
	if db.Error != nil {
		return db.Error
	}
	if key == nil {
		return fmt.Errorf("invalid api key: %s", apiKey)
	}
	if key.UsedNum >= key.Limit {
		return fmt.Errorf("Available times:%d, has used times: %d", key.Limit, key.UsedNum)
	}
	return nil
}

func (this *SagaDB) Close() error {
	return this.db.Close()
}
