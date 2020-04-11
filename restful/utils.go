package restful

import (
	"github.com/qiangxue/fasthttp-routing"
	"encoding/json"
	"fmt"
)

func GetParam(ctx *routing.Context, params ...string) (map[string]string, int64) {
	res := make(map[string]string)
	for _, param := range params {
		res[param] = ctx.Param(param)
		if res[param] == "" {
			return nil, PARA_ERROR
		}
	}
	return res, SUCCESS
}

func PostParam(ctx *routing.Context) (map[string]interface{}, error) {
	req := ctx.PostBody()
	if req == nil || len(req) == 0 {
		return nil, fmt.Errorf("param length is 0")
	}
	arg := make(map[string]interface{})
	err := json.Unmarshal(req, &arg)
	if err != nil {
		return nil, fmt.Errorf("param Unmarshal error: %s", err)
	}
	param, ok := arg["params"]
	if !ok {
		return nil, fmt.Errorf("arg[params] is not ok")
	}
	para, ok := param.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("param.(map[string]interface{})")
	}
	return para, nil
}