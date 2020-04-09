package nasa

import (
	"fmt"
	"github.com/ontio/saga/config"
	"github.com/ontio/saga/core/http"
	"github.com/ontio/saga/dao"
)

var (
	apod = "https://api.nasa.gov/planetary/apod?api_key=%s"
	feed = "https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&end_date=%s&api_key=%s"
)

type Nasa struct {
	requestNum int
	db         *dao.NasaDB
}

func NewNasa() (*Nasa, error) {
	db, err := dao.NewNasaDB()
	if err != nil {
		return nil, err
	}
	num, err := db.QueryRequestNum()
	if err != nil {
		return nil, err
	}
	return &Nasa{
		requestNum: num,
		db:         db,
	}, nil
}

func (this *Nasa) Apod() ([]byte, error) {
	url := fmt.Sprintf(apod, config.NASA_API_KEY)
	return http.Get(url)
}

func (this *Nasa) Feed(startDate, endDate string) ([]byte, error) {
	url := fmt.Sprintf(feed, startDate, endDate, config.NASA_API_KEY)
	return http.Get(url)
}
