package restful

import "github.com/qiangxue/fasthttp-routing"

const (
	POST_EXECL = "/api/v1/nasa/apod"
)

const (
	GET_NASA_APOD = "/api/v1/nasa/apod"
	GET_NASA_FEED = "/api/v1/nasa/feed/<startdate>/<enddate>"
)

//init restful server
func InitRouter() *routing.Router {
	router := routing.New()
	//router.Post(POST_EXECL, nil)
	router.Get(GET_NASA_APOD, Apod)
	router.Get(GET_NASA_APOD, Feed)
	return router
}
