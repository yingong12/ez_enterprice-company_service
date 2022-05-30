package controller

import (
	"company_service/http/buz_code"
	"company_service/http/request"
	"company_service/http/response"
	"company_service/service"
	"fmt"
	"net/http"

	"company_service/logger"

	"github.com/gin-gonic/gin"
)

//GetProductInfo 企业搜索
//@Summary	企业搜索
//@Description	企业搜索
//@Tags	常规接口
//@Produce	json
//@Param	xxx query request.ProductInfoRequest  false "字段注解"
//@Success 200 {object} response.ProductInfo
//@Router	/data_analysis/common/get_product_info [post]
func Search(ctx *gin.Context) {
	req := request.Search{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	list, total, err := service.Search(req.RangeFilters, req.TextFilters, req.Sort, req.Page, req.PageSize)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	buzCode := buz_code.CODE_OK
	msg := "ok"
	data := response.Search{
		List:  list,
		Total: total,
	}

	ctx.JSON(http.StatusOK, gin.H{"code": buzCode, "msg": msg, "data": data})
}
