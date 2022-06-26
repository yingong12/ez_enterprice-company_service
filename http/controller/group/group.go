package group

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/group"
	service "company_service/service/group"
	"log"
	"strings"

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
	appIDs := strings.Split(req.AppIDs, ",")
	log.Println(appIDs, len(appIDs), appIDs[0] == "")
	list, total, err := service.Search(appIDs, req.Name, req.Sort, req.Page, req.PageSize)
	data := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	res.Data = data
	//
	return
}
func GetChildrenMulti(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.GetChildrenMulti{}
	err = BindQuery(ctx, &req)
	if err != nil {
		return
	}
	//
	return
}
func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := group.Create{}
	err = BindJSON(ctx, &req)
	if err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	err = service.Create(req.UID, req.Data)
	if err != nil {
		res.Code = buz_code.CODE_ENTERPRISE_CREATE_FAILED
		res.Msg = err.Error()
	}
	return
}
func Update(ctx *gin.Context) (res controller.STDResponse, err error) {
	return
}
