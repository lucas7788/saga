package restful

import "github.com/qiangxue/fasthttp-routing"

const (
	POST_EXECL = "/api/v1/uploadexecl"
)

const (
	GET_EVENT_TYPE = "/api/v1/getevtty"
)

//init restful server
func InitRouter() *routing.Router {
	router := routing.New()
	router.Post(POST_EXECL, nil)
	router.Get(GET_EVENT_TYPE, nil)
	return router
}
