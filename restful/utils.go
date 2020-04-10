package restful

import "github.com/qiangxue/fasthttp-routing"

func GetParam(ctx *routing.Context, params ...string) map[string]interface{} {
	res := make(map[string]interface{})
	for _, param := range params {
		res[param] = ctx.Param(param)
	}
	return res
}
