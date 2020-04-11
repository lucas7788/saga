package restful

import (
	"encoding/json"
	"fmt"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/saga/dao"
	"github.com/qiangxue/fasthttp-routing"
	"sync"
)

const (
	POST_PAY = "/api/v1/pay"
)
const (
	GET_SEARCH_API   = "/api/v1/searchapi/<key>/<apikey>"
	GET_NASA_APOD    = "/api/v1/nasa/apod/<apikey>"
	GET_NASA_FEED    = "/api/v1/nasa/feed/<startdate>/<enddate>/<apikey>"
	GET_ALL_API_INFO = "/api/v1/getallapiinfo"
)

var getMethodMap, postMethodMap map[string]Action

//init restful server
func InitRouter() *routing.Router {
	router := routing.New()
	registryMethod()
	for path, v := range getMethodMap {
		router.Get(path, func(context *routing.Context) error {
			req, errCode := getParam(context, path)
			if errCode != SUCCESS {
				return writeResponse(context, ResponsePack(errCode, nil))
			}
			err := verifyApiKey(path, req["apikey"])
			var resp map[string]interface{}
			if err != nil {
				log.Errorf("parse get request param error: %s", err)
				resp = ResponsePack(INTER_ERROR, err)
			} else {
				resp = v.handler(req)
			}
			resp["Action"] = v.name
			return writeResponse(context, resp)
		})
	}

	for path, v := range postMethodMap {
		router.Post(path, func(context *routing.Context) error {
			reqParam, err := PostParam(context)
			var resp map[string]interface{}
			if err != nil {
				log.Errorf("parse post request param error: %s", err)
				resp = ResponsePack(INTER_ERROR, err)
			} else {
				resp = v.handler(reqParam)
			}
			resp["Action"] = v.name
			return writeResponse(context, resp)
		})
	}
	return router
}

func getParam(context *routing.Context, url string) (map[string]string, int64) {
	reqParam := make(map[string]string)
	var errCode int64
	switch url {
	case GET_NASA_APOD:
		reqParam, errCode = GetParam(context, "apikey")
	case GET_NASA_FEED:
		reqParam, errCode = GetParam(context, "startdate", "endate", "apikey")
	case GET_SEARCH_API:
		reqParam, errCode = GetParam(context, "key", "apikey")
	}
	return reqParam, errCode
}

func verifyApiKey(url string, apiKey interface{}) error {
	switch url {
	case GET_SEARCH_API, POST_PAY:
		return nil
	}
	key, ok := apiKey.(string)
	if !ok {
		return fmt.Errorf("error apikey: %v", apiKey)
	}
	return dao.DefDB.VerifyApiKey(key)
}

func writeResponse(ctx *routing.Context, res interface{}) error {
	bs, err := json.Marshal(res)
	if err != nil {
		return err
	}
	l, err := ctx.Write(bs)
	if l != len(bs) || err != nil {
		log.Errorf("write error: %s, expected length: %d, actual length: %d", err, len(bs), l)
		return err
	}
	return nil
}

type handler func(map[string]string) map[string]interface{}

type Action struct {
	sync.RWMutex
	name    string
	handler handler
}

func registryMethod() {
	getMethodMap = map[string]Action{
		GET_NASA_APOD:    {name: "apod", handler: Apod},
		GET_NASA_FEED:    {name: "feed", handler: Feed},
		GET_SEARCH_API:   {name: "searchapi", handler: SearchApi},
		GET_ALL_API_INFO: {name: "getallapiinfo"},
	}

	postMethodMap = map[string]Action{
		POST_PAY: {name: "pay", handler: Pay},
	}
}
