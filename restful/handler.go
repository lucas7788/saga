package restful

import (
	"github.com/ontio/saga/config"
	"github.com/ontio/saga/core/nasa"
	"github.com/ontio/saga/dao"
	"sync"
)

var DefMap = new(sync.Map)

func SearchApi(params map[string]interface{}) map[string]interface{} {
	key, _ := params["key"].(string)
	dao.DefDB.SearchApi(key)
	return nil
}

func Apod(params map[string]interface{}) map[string]interface{} {
	nas, err := getNasa()
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	res, err := nas.Apod()
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}

func Feed(params map[string]interface{}) map[string]interface{} {
	//TODO param check
	nas, err := getNasa()
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	startDate, _ := params["startdate"].(string)
	endDate, _ := params["enddate"].(string)
	res, err := nas.Feed(startDate, endDate)
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}

func getNasa() (*nasa.Nasa, error) {
	var nas *nasa.Nasa
	val, ok := DefMap.Load(config.NASA_NAME)
	if !ok || val == nil {
		var err error
		nas, err = nasa.NewNasa()
		if err != nil {
			return nil, err
		}
		DefMap.Store(config.NASA_NAME, nas)
	} else {
		nas = val.(*nasa.Nasa)
	}
	return nas, nil
}
