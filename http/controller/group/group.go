package group

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/group"
	"company_service/model"
	"company_service/providers"
	service "company_service/service/group"
	"company_service/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindJSON(ctx *gin.Context, req interface{}) (err error) {
	err = ctx.BindJSON(req)
	return
}
func BindQuery(ctx *gin.Context, form interface{}) (err error) {
	err = ctx.BindQuery(form)
	return
}
func BindMultiForm(ctx *gin.Context, form interface{}) (err error) {
	err = ctx.BindWith(form, binding.FormMultipart)
	return
}

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
	url := fmt.Sprintf("%saccount/init_app?app_type=1&uid=%s&b_access_token=%s", client.BaseURL, req.UID, ctx.GetHeader("b_access_token"))
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
		err = errors.New("转发返回值有误")
		return
	}
	appID := httpRes["data"].(string)
	//
	//
	data := model.GroupMuttable{
		Name:               utils.GenStringWithPrefix("机构_grp", 16),
		RegistrationNumber: utils.GenStringWithPrefix("信用代码", 14),
	}
	err = service.Create(appID, req.UID, data)
	res.Data = appID
	return
}
func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Search{}
	err = BindQuery(ctx, &req)
	if err != nil {
		return
	}
	list, total, err := service.Search(req.AppID, req.Name, req.Sort, req.Page, req.PageSize)
	for k := range list {
		if d := utils.DFSDistrict(&providers.DisrictDict, list[k].District); d != nil {
			list[k].LabelDistrict = []string{d.Label}
		}
		if d := utils.DFSIndustry(&providers.IndustryDict, list[k].Industry); d != nil {
			list[k].LabelIndustry = []string{d.Label}
		}
	}
	data := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	res.Data = data
	//
	return
}

//批量查询儿子企业信息
func GetChildrenMulti(ctx *gin.Context) (res controller.STDResponse, err error) {
	//
	req := group.GetChildrenMulti{}
	err = BindQuery(ctx, &req)
	if err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		return
	}
	list, total, err := service.ChilrenInfo(req.AppID, req.Page, req.PageSize)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		return
	}
	res.Data = map[string]interface{}{
		"list":  list,
		"total": total,
	}
	//
	return
}
func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Create{}
	err = BindJSON(ctx, &req)
	if err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		return
	}
	err = service.Create("", req.UID, req.Data)
	if err != nil {
		res.Code = buz_code.CODE_ENTERPRISE_CREATE_FAILED
	}
	return
}

//TODO:
func Update(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Update{}
	appID := ctx.Param("app_id")
	err = BindJSON(ctx, &req)
	if err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		return
	}
	rf, err := service.Update(appID, req.Data, -1)
	if err != nil {
		res.Code = buz_code.CODE_ENTERPRISE_UPDATE_FAILED
	}
	res.Data = map[string]interface{}{
		"affected_rows": rf,
	}
	return
}
