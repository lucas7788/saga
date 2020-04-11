package models

import (
	"github.com/ontio/saga/config"
	"time"
)

type BuyRecord struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	OntId     string
	UserName  string
	TxHash    string
	Price     string
	ApiId     string
	ApiKey    string
	TxStatus  config.TxStatus
}

type APIInfo struct {
	ID       uint `gorm:"primary_key"`
	ApiId    int
	ApiUrl   string
	ApiPrice string
	ApiDesc  string
}

type APIKey struct {
	ID       uint `gorm:"primary_key"`
	ApiKey   string
	ApiId    int
	Limit    int
	UsedNum  int
	OntId    string
	UserName string
}
