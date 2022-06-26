package group

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/group"
	service "company_service/service/group"

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

func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Search{}
	err = BindQuery(ctx, &req)
	if err != nil {
		return
	}
	list, total, err := service.Search(req.AppID, req.Name, req.Sort, req.Page, req.PageSize)
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
	err = service.Create(req.UID, req.Data)
	if err != nil {
		res.Code = buz_code.CODE_ENTERPRISE_CREATE_FAILED
	}
	return
}
func Update(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Update{}
	appID := ctx.Param("app_id")
	err = BindJSON(ctx, &req)
	if err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		return
	}
	rf, err := service.Update(appID, req.GroupMuttable)
	if err != nil {
		res.Code = buz_code.CODE_ENTERPRISE_UPDATE_FAILED
	}
	res.Data = map[string]interface{}{
		"affected_rows": rf,
	}
	return
}
