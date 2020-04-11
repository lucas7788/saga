package restful

import (
	"github.com/ontio/saga/core/nasa"
	"github.com/ontio/saga/dao"
	"strconv"
)

func GetApiInfoByPage(params map[string]interface{}) map[string]interface{} {
	pageNumStr, _ := params["pageNum"].(string)
	pageSizeStr, _ := params["pageSize"].(string)
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		return ResponsePack(PARA_ERROR, err)
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
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

func SearchApi(params map[string]interface{}) map[string]interface{} {
	key, ok := params["key"].(string)
	if !ok || key == "" {
		return ResponsePack(PARA_ERROR, nil)
	}
	infos, err := dao.DefDB.SearchApi(key)
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(infos)
}

func Apod(params map[string]interface{}) map[string]interface{} {
	res, err := nasa.Apod()
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}

func Feed(params map[string]interface{}) map[string]interface{} {
	//TODO param check
	startDate, ok := params["startdate"].(string)
	if !ok || startDate == "" {
		return ResponsePack(PARA_ERROR, nil)
	}
	endDate, ok := params["enddate"].(string)
	if !ok || endDate == "" {
		return ResponsePack(PARA_ERROR, nil)
	}
	res, err := nasa.Feed(startDate, endDate)
	if err != nil {
		return ResponsePack(INTER_ERROR, err)
	}
	return ResponseSuccess(res)
}
