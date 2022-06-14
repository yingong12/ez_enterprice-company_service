package enterprise

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request"
	"company_service/http/response"
	"company_service/model"
	"company_service/providers"
	"company_service/service"
	"company_service/utils"
	"log"
	"strings"

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
func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := request.Search{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	list, total, err := service.Search(req.RangeFilters, req.TextFilters, req.Sort, req.Page, req.PageSize)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	data := response.Search{
		List:  list,
		Total: total,
	}
	res.Data = data
	return
}

func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := request.Create{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	buzCode, msg, err := service.Create(req.UID, req.ParentID, req.Data)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Code = buzCode
	res.Msg = msg
	return
}

func Update(ctx *gin.Context) (res controller.STDResponse, err error) {
	appID, ok := ctx.Params.Get("app_id")
	if !ok {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	req := request.Update{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	rows, err := service.Update(appID, req.Data)
	if err != nil {
		//TODO:几种业务报错怎么更优雅的去弄
		//把具体哪个键报出来
		if utils.IsMysqlDupKeyErr(err) {
			res.Code = buz_code.CODE_ENTERPRISE_UPDATE_FAILED
			res.Msg = err.Error()
			return
		}
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	if rows <= 0 {
		res.Code = buz_code.CODE_ENTERPRISE_UPDATE_FAILED
		res.Msg = "没有该企业"
		return
	}
	return
}
func QueryByIDs(ctx *gin.Context) (res controller.STDResponse, err error) {
	//单次最多50个
	idSlice := strings.Split(ctx.Query("app_ids"), ",")
	if len(idSlice) > 50 {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = "最多查询50个企业"
		return
	}
	list, err := service.GetByAppIDs(idSlice)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Data = list
	return
}

func GetIndustryByCode(ctx *gin.Context) (res controller.STDResponse, err error) {
	//DFS
	//根据ID拿节点以及儿子
	code := ctx.Query("code")
	node := dfsIndustry(&providers.IndustryDict, code)
	if node == nil {
		res.Code = buz_code.CODE_NODE_NOT_FOUND
		res.Msg = "节点不存在"
		return
	}
	children := []*model.IndustryDict{}
	for _, v := range node.Children {
		item := &model.IndustryDict{
			Code:     v.Code,
			Label:    v.Label,
			IsLeaf:   v.Children == nil,
			Children: nil,
		}
		children = append(children, item)
	}
	data := model.IndustryDict{
		Children: children,
		Code:     node.Code,
		Label:    node.Label,
		IsLeaf:   len(children) == 0,
	}
	res.Data = data
	return
}

func GetDistrictByCode(ctx *gin.Context) (res controller.STDResponse, err error) {
	//DFS
	code := ctx.Query("code")
	log.Println(code)
	node := dfsDistrict(&providers.DisrictDict, code)
	if node == nil {
		res.Code = buz_code.CODE_NODE_NOT_FOUND
		res.Msg = "节点不存在"
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
		IsLeaf:   len(children) == 0,
	}
	res.Data = data
	return
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
