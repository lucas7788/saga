package restful

import "github.com/qiangxue/fasthttp-routing"

func GetParam(ctx *routing.Context, params ...string) []string {
	res := make([]string, len(params))
	for index,param := range params {
		res[index] = ctx.Param(param)
	}
	return res
}