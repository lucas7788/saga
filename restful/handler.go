package restful

import (
	"github.com/ontio/saga/core/nasa"
	"github.com/ontio/saga/dao"
	"strconv"
)

func GetApiInfoByPage(params map[string]string) map[string]interface{} {
	pageNumStr,_ := params["pageNum"]
	pageSizeStr,_ := params["pageSize"]
	pageNum,err := strconv.Atoi(pageNumStr)
	if err != nil {
		return ResponsePack(PARA_ERROR, err)
	}
	pageSize,err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return ResponsePack(PARA_ERROR, err)
	}
	if pageNum < 1 {
		pageNum = 1
	}
	start := (pageNum - 1) * pageSize
	res, err := dao.DefDB.QueryApiInfoByPage(start, pageSize)
	if err != nil {
		return ResponsePack(SQL_ERROR, err)
	}
	return ResponseSuccess(res)
}

func SearchApi(params map[string]string) map[string]interface{} {
	key, ok := params["key"]
	if !ok || key == "" {
		return ResponsePack(PARA_ERROR, nil)
	}
	infos,err := dao.DefDB.SearchApi(key)
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(infos)
}

func Pay(params map[string]interface{}) map[string]interface{} {

	return nil
}

func Apod(params map[string]string) map[string]interface{} {
	res, err := nasa.Apod()
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}

func Feed(params map[string]string) map[string]interface{} {
	//TODO param check
	startDate, ok := params["startdate"]
	if !ok || startDate == ""{
		return ResponsePack(PARA_ERROR, nil)
	}
	endDate, ok := params["enddate"]
	if !ok || endDate == "" {
		return ResponsePack(PARA_ERROR, nil)
	}
	res, err := nasa.Feed(startDate, endDate)
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}
