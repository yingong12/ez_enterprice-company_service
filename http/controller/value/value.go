package value

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/valuate"
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
func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := valuate.Search{}
	if err = BindQuery(ctx, &req); err != nil {
		return
	}
	list, err := service.Search(req.AppID, req.Page, req.PageSize)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Data = list
	return
}

//提交估值
func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := valuate.Create{}
	if err = BindJSON(ctx, &req); err != nil {
		return
	}
	err = service.Create(req.ValuateMuttable)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	return
}

func Export(ctx *gin.Context) {}
