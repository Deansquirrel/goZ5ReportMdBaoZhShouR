package router

import (
	"encoding/json"
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
	"io/ioutil"
)

const (
	TranErrStr = "{\"errcode\":-1,\"errmsg\":\"构造返回结果时发生错误, %s\"}"
)

type common struct {
}

func (c *common) GetRequestBody(ctx iris.Context) string {
	body := ctx.Request().Body
	defer func() {
		_ = body.Close()
	}()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error("获取Http请求文本时发生错误：" + err.Error())
		return ""
	}
	return string(b)
}

//向ctx中添加返回内容
func (c *common) WriteResponse(ctx iris.Context, v interface{}) {
	str, err := json.Marshal(v)
	if err != nil {
		body := fmt.Sprintf(TranErrStr, "err:"+err.Error())
		_, err = ctx.WriteString(body)
		if err != nil {
			log.Error(fmt.Sprintf("write body err,body: %s,err: %s", string(str), err.Error()))
		}
		return
	}
	_, err = ctx.WriteString(string(str))
	if err != nil {
		log.Error(fmt.Sprintf("write body err,body: %s,err: %s", string(str), err.Error()))
	}
	return
}
