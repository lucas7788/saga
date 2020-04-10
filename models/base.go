package models

import "time"

type BuyRecord struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	OntId     string
	UserName  string
	TxHash    string
	Price     string
	APIId     string
	ApiKey    string
}

type APIInfo struct {
	ID       uint `gorm:"primary_key"`
	APIId    int
	APIUrl   string
	APIPrice string
	APIDesc  string
}

type APIKey struct {
	ID       uint `gorm:"primary_key"`
	ApiKey   string
	APIId    int
	Limit    int
	UsedNum  int
	OntId    string
	UserName string
}
