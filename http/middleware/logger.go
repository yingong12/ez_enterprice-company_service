package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//记录controller抛出的错误
func ControllerErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err, ok := ctx.Get("buz_err")
		if err == nil || !ok {
			return
		}
		log.Printf("[endpoint]:%s [error]:%v\n", ctx.Request.URL.Path, err)
	}
}

//访问日志
func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//后置
		before := time.Now()
		bodyStream := copyBody(ctx)
		ctx.Next()

		after := time.Now()
		//配置到单独logger
		log.Printf("[response_time]:%dms, [end_point]:%s, [method]:%s, [body]:%s\n", after.Sub(before).Milliseconds(), ctx.Request.URL.Path, ctx.Request.Method, removeJSONIndent(bodyStream))
	}
}

func removeJSONIndent(input []byte) (output []byte) {
	bstd := map[string]interface{}{}
	json.Unmarshal(input, &bstd)
	output, _ = json.Marshal(bstd)
	return
}

//copy request body
func copyBody(ctx *gin.Context) (buf []byte) {
	dist := &bytes.Buffer{}
	//从src读写到dst
	trdr := io.TeeReader(ctx.Request.Body, dist)
	//注意 这里err了请求就丢了
	buf, _ = ioutil.ReadAll(trdr)
	ctx.Request.Body = io.NopCloser(dist)
	return
}
