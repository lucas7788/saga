package restful

import "github.com/qiangxue/fasthttp-routing"

func ParseFeedParam(ctx *routing.Context) (startDate, endDate string) {
	startDate = ctx.Param("startdate")
	endDate = ctx.Param("enddate")
	return
}
