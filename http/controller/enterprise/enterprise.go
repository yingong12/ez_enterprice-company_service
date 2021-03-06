package enterprise

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request"
	"company_service/http/response"
	"company_service/model"
	"company_service/providers"
	en "company_service/repository"
	"company_service/service"
	"company_service/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestInit struct {
	UID string `json:"uid"`
}

func Init(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := &RequestInit{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	//
	client := providers.HttpClientAccount
	url := fmt.Sprintf("%saccount/init_app?uid=%s&b_access_token=%s", client.BaseURL, req.UID, ctx.GetHeader("b_access_token"))
	rsp, err := client.Get(url)
	if err != nil || rsp.StatusCode != 200 {
		log.Println(err, 37, rsp.StatusCode)
		res.Code = buz_code.CODE_SERVER_ERROR
		return
	}
	httpRes := map[string]interface{}{}
	buf, err := ioutil.ReadAll(rsp.Body)
	log.Println((string)(buf), 42)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		return
	}
	err = json.Unmarshal(buf, &httpRes)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		return
	}
	if httpRes["code"].(float64) != 0 {
		res.Code = buz_code.CODE_SERVER_ERROR
		err = errors.New(httpRes["msg"].(string))
		return
	}
	appID := httpRes["data"].(string)
	//
	//
	data := model.EnterpriseMuttable{
		Name:               utils.GenStringWithPrefix("企业_app", 16),
		RegistrationNumber: utils.GenStringWithPrefix("信用代码", 14),
	}
	err = en.Create(appID, req.UID, "", data)
	res.Data = appID
	return

}
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
	//添加label, 行业，地区。
	for k := range list {
		v := &list[k]
		codeInd := v.Industry
		codeDist := v.District
		if path := getPathDistrict(&providers.DisrictDict, codeDist); len(path) > 0 {
			v.LabelDistrict = path[1:]
		}
		if path := getPathIndustry(&providers.IndustryDict, codeInd); len(path) > 0 {
			v.LabelIndustry = path[1:]
		}
	}
	data := response.Search{
		List:  list,
		Total: total,
	}
	res.Data = data
	return
}

//GetProductInfo 企业新建
//@Summary	企业新建
//@Description	企业新建
//@Tags	企业
//@Produce	json
//@Param	xxx body request.Create  false "字段注解"
//@Router	/enterprise [POST]
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
	node := utils.DFSIndustry(&providers.IndustryDict, code)
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
	node := utils.DFSDistrict(&providers.DisrictDict, code)
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

//DFS top-down 带入当前路劲。 bottom-up
func getPathDistrict(root *model.District, target string) (path []string) {
	if root == nil {
		return
	}
	if root.Code == target {
		path = []string{root.Label}
		return
	}
	for _, d := range root.Children {
		if cur := getPathDistrict(d, target); len(cur) > 0 {
			path = append([]string{root.Label}, cur...)
			return
		}
	}
	return
}
func getPathIndustry(root *model.IndustryDict, target string) (path []string) {
	if root == nil {
		return
	}
	if root.Code == target {
		path = []string{root.Label}
		return
	}
	for _, d := range root.Children {
		if cur := getPathIndustry(d, target); len(cur) > 0 {
			path = append([]string{root.Label}, cur...)
			return
		}
	}
	return
}
