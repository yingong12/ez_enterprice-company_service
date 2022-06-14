package controller

import "github.com/gin-gonic/gin"

func Health(ctx *gin.Context) {
	ctx.Writer.Write([]byte("ok"))
}
