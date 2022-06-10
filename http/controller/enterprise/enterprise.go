package enterprise

import (
	"company_service/http/buz_code"
	"company_service/http/request"
	"company_service/http/response"
	"company_service/model"
	"company_service/providers"
	"company_service/service"
	"company_service/utils"
	"fmt"
	"log"
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
	code := ctx.Query("code")
	node := dfsIndustry(&providers.IndustryDict, code)
	if node == nil {
		return
	}
	res := []*model.IndustryDict{}
	for _, v := range node.Children {
		item := &model.IndustryDict{
			Code:     v.Code,
			Label:    v.Label,
			IsLeaf:   v.Children == nil,
			Children: nil,
		}
		res = append(res, item)
	}
	data := model.IndustryDict{
		Children: res,
		Code:     node.Code,
		Label:    node.Label,
		IsLeaf:   len(res) == 0,
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": data})
}

func GetDistrictByCode(ctx *gin.Context) {
	//DFS
	code := ctx.Query("code")
	log.Println(code)
	node := dfsDistrict(&providers.DisrictDict, code)
	if node == nil {
		return
	}
	children := []*model.District{}
	for _, v := range node.Children {
		//删除children
		//这里克隆一份防止改动原有全局数据
		item := &model.District{
			Code:     v.Code,
			Label:    v.Label,
			Level:    v.Level,
			IsLeaf:   v.Children == nil,
			Children: nil,
		}
		children = append(children, item)
	}
	data := model.District{
		Children: children,
		Code:     node.Code,
		Label:    node.Label,
		Level:    node.Level,
		IsLeaf:   children == nil,
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": data})
}

//找出该节点和所有儿子
func dfsDistrict(root *model.District, target string) *model.District {
	if root == nil {
		return nil
	}
	if root.Code == target {
		return root
	}
	for _, d := range root.Children {
		if cur := dfsDistrict(d, target); cur != nil {
			return cur
		}
	}
	return nil
}
func dfsIndustry(root *model.IndustryDict, target string) *model.IndustryDict {
	if root == nil {
		return nil
	}
	if root.Code == target {
		return root
	}
	for _, d := range root.Children {
		if cur := dfsIndustry(d, target); cur != nil {
			return cur
		}
	}
	return nil
}
