package restful

import (
	"encoding/json"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/saga/config"
	"github.com/ontio/saga/core/nasa"
	"github.com/qiangxue/fasthttp-routing"
	"sync"
)

var DefMap = new(sync.Map)

func Apod(ctx *routing.Context) error {
	nas, err := getNasa()
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	res, err := nas.Apod()
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	return writeResponse(ctx, ResponseSuccess(res))
}

func Feed(ctx *routing.Context) error {
	params := GetParam(ctx, "startdate", "endate")
	//TODO param check
	nas, err := getNasa()
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	res, err := nas.Feed(params[0], params[1])
	if err != nil {
		return writeResponse(ctx, ResponsePack(INTER_ERROR, err))
	}
	return writeResponse(ctx, ResponseSuccess(res))
}

func getNasa() (*nasa.Nasa, error) {
	var nas *nasa.Nasa
	val, ok := DefMap.Load(config.NASA_NAME)
	if !ok || val == nil {
		var err error
		nas, err = nasa.NewNasa()
		if err != nil {
			return nil, err
		}
		DefMap.Store(config.NASA_NAME, nas)
	} else {
		nas = val.(*nasa.Nasa)
	}
	return nas, nil
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
