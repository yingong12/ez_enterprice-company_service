package enterprise

import (
	"company_service/http/buz_code"
	"company_service/http/request"
	"company_service/http/response"
	"company_service/service"
	"company_service/utils"
	"fmt"
	"net/http"
	"strings"

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

func Create(ctx *gin.Context) {
	req := request.Create{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	buzCode, msg, err := service.Create(req.UID, req.ParentID, req.Data)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buzCode, "msg": msg, "data": ""})
}

func Update(ctx *gin.Context) {
	appID, ok := ctx.Params.Get("app_id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", "no app_id"),
		})
		return
	}
	req := request.Update{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	rows, err := service.Update(appID, req.Data)
	if err != nil {
		//TODO:几种业务报错怎么更优雅的去弄
		//把具体哪个键报出来
		if utils.IsMysqlDupKeyErr(err) {
			ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_ENTERPRISE_UPDATE_FAILED, "msg": "唯一键冲突", "data": ""})
			return
		}
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	if rows <= 0 {
		ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_ENTERPRISE_UPDATE_FAILED, "msg": "没有该企业", "data": ""})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": ""})
}
func QueryByIDs(ctx *gin.Context) {
	//单次最多50个
	idSlice := strings.Split(ctx.Query("app_ids"), ",")
	if len(idSlice) > 50 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("最多查询50个"),
		})
		return
	}
	list, err := service.GetByAppIDs(idSlice)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": list})
}

func GetIndustryByCode(ctx *gin.Context) {
	//DFS
	//根据ID拿节点以及儿子
}

func GetDistrictByCode(ctx *gin.Context) {
	//DFS
}
