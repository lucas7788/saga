package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type NasaDB struct {
	db *sql.DB
}

func NewNasaDB() (*NasaDB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	return &NasaDB{
		db: db,
	}, nil
}

func (this *NasaDB) UpdateRequestNum(num int) {
	this.db.Exec("")
}

func (this *NasaDB) QueryRequestNum() (int, error) {
	return 0, nil
}

func (this *NasaDB) CloseDB() error {
	return this.db.Close()
}
