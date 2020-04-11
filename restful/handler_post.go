package restful

import (
	"github.com/ontio/saga/dao"
	"strconv"
)

func ApiTest(params map[string]interface{}) map[string]interface{} {
	apiIdStr, ok := params["apiId"].(string)
	if !ok {
		return ResponsePack(PARA_ERROR, nil)
	}
	//ontid,ok := params["ontid"].(string)
	//if !ok {
	//	return ResponsePack(PARA_ERROR, nil)
	//}
	//useName,ok := params["useName"].(string)
	//if !ok {
	//	return ResponsePack(PARA_ERROR, nil)
	//}
	apiId, err := strconv.ParseUint(apiIdStr, 10, 64)
	if err != nil {
		return ResponsePack(PARA_ERROR, err)
	}
	dao.DefDB.QueryApiInfoByApiId(uint(apiId))
	return nil
}

func Pay(params map[string]interface{}) map[string]interface{} {

	return nil
}
