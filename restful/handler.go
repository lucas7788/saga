package restful

import (
	"github.com/qiangxue/fasthttp-routing"
	"encoding/json"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/saga/core/nasa"
)

func Apod(ctx *routing.Context) error {
	res, err := nasa.Apod()
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	return writeResponse(ctx, ResponseSuccess(res))
}

func Feed(ctx *routing.Context) error{
	startDate, endDate := ParseFeedParam(ctx)
	//TODO param check
	res, err := nasa.Feed(startDate, endDate)
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	return writeResponse(ctx, ResponseSuccess(res))
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
