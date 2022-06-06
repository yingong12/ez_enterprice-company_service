package value

import (
	"company_service/http/buz_code"
	"company_service/http/request/valuate"
	"company_service/logger"
	service "company_service/service/valuate"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(ctx *gin.Context, req interface{}) (err error) {
	if err = ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
	}
	return
}
func BindQuery(ctx *gin.Context, form interface{}) (err error) {
	if err = ctx.BindQuery(form); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
	}
	return
}

//搜索接口
func Search(ctx *gin.Context) {
	req := valuate.Search{}
	if err := BindQuery(ctx, &req); err != nil {
		return
	}
	list, err := service.Search(req.AppID, req.Page, req.PageSize)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": buz_code.CODE_OK,
		"msg":  "ok",
		"list": list,
	})
}

//提交估值
func Create(ctx *gin.Context) {
	req := valuate.Create{}
	if err := BindJSON(ctx, &req); err != nil {
		return
	}
	err := service.Create(req.ValuateMuttable)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": buz_code.CODE_OK,
		"msg":  "ok",
	})
}

func Export(ctx *gin.Context) {}
