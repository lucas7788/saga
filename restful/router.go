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
	GET_SEARCH_API = "/api/v1/searchapi/<key>/<apikey>"
	GET_NASA_APOD  = "/api/v1/nasa/apod/<apikey>"
	GET_NASA_FEED  = "/api/v1/nasa/feed/<startdate>/<enddate>/<apikey>"
)

var getMethodMap map[string]Action

//init restful server
func InitRouter() *routing.Router {
	router := routing.New()
	registryMethod()
	for path, v := range getMethodMap {
		router.Get(path, func(context *routing.Context) error {
			req := getParam(context, path)
			err := verifyApiKey(path, req["apikey"])
			var resp map[string]interface{}
			if err != nil {
				resp = ResponsePack(INTER_ERROR, err)
			} else {
				resp = v.handler(req)
			}
			resp["Action"] = v.name
			return writeResponse(context, resp)
		})
	}
	return router
}

func getParam(context *routing.Context, url string) map[string]interface{} {
	reqParam := make(map[string]interface{})
	switch url {
	case GET_NASA_APOD:
		reqParam = GetParam(context, "apikey")
	case GET_NASA_FEED:
		reqParam = GetParam(context, "startdate", "endate", "apikey")
	case GET_SEARCH_API:
		reqParam = GetParam(context, "key", "apikey")
	}
	return reqParam
}
func verifyApiKey(url string, apiKey interface{}) error {
	switch url {
	case GET_SEARCH_API:
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

type handler func(map[string]interface{}) map[string]interface{}

type Action struct {
	sync.RWMutex
	name    string
	handler handler
}

func registryMethod() {
	getMethodMap = map[string]Action{
		GET_NASA_APOD:  {name: "apod", handler: Apod},
		GET_NASA_FEED:  {name: "feed", handler: Feed},
		GET_SEARCH_API: {name: "searchapi", handler: SearchApi},
	}
}
