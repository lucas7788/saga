package restful

import (
	"github.com/qiangxue/fasthttp-routing"
)

const (
	POST_EXECL = "/api/v1/nasa/apod"
)

const (
	GET_NASA_APOD = "/api/v1/nasa/apod"
	GET_NASA_FEED = "/api/v1/nasa/feed/<startdate>/<enddate>"
)

var getMethodMap map[string]routing.Handler
//init restful server
func InitRouter() *routing.Router {
	router := routing.New()
	registryMethod()
	for k,v := range getMethodMap {
		router.Get(k, v)
	}
	return router
}
func registryMethod() {
	getMethodMap = map[string]routing.Handler{
		GET_NASA_APOD: Apod,
		GET_NASA_FEED: Feed,
	}
}
