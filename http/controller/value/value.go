package value

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/valuate"
	"company_service/model"
	service "company_service/service/valuate"
	"fmt"
	"net/http"

	_ "company_service/model"

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

//GetProductInfo 估值搜索
//@Summary	估值搜索
//@Description	估值搜索
//@Tags 估值
//@Produce	json
//@Param	xxx query valuate.Search  false "字段注解"
//@Success 200 {object} []model.Valuate
//@Router	/valuate/search [get]
func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := valuate.Search{}
	if err = BindQuery(ctx, &req); err != nil {
		return
	}
	list, total, err := service.Search(req.AppID, req.Page, req.PageSize)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	if list == nil {
		list = []model.Valuate{}
	}
	res.Data = map[string]interface{}{
		"list":  list,
		"total": total,
	}
	return
}

//提交估值
func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := valuate.Create{}
	if err = BindJSON(ctx, &req); err != nil {
		return
	}
	err = service.Create(req)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	return
}

func Export(ctx *gin.Context) {}
