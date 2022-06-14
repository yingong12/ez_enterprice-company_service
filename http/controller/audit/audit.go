package audit

import (
	"company_service/http/buz_code"
	"company_service/http/controller"
	"company_service/http/request/audit"
	"company_service/http/response"
	service "company_service/service/audit"
	"company_service/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func BindJSON(ctx *gin.Context, req interface{}) (err error) {
	ctx.BindJSON(req)
	return
}
func BindQuery(ctx *gin.Context, form interface{}) (err error) {
	ctx.BindQuery(form)
	return
}
func BindMultiForm(ctx *gin.Context, form interface{}) (err error) {
	ctx.BindWith(form, binding.FormMultipart)
	return
}
func Search(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := audit.Search{}
	if err = BindQuery(ctx, &req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	stateArr := []int{}
	if req.States != "" {
		//bind states
		stateStrArr := strings.Split(req.States, ",")
		for _, state := range stateStrArr {
			stateInt, errInt := strconv.Atoi(state)
			if errInt != nil {
				err = errInt
				res.Code = buz_code.CODE_INVALID_ARGS
				res.Msg = err.Error()
				return
			}
			stateArr = append(stateArr, stateInt)
		}
	}

	list, count, err := service.Search(req.AppName, req.RegistrationNumber, req.AppID, stateArr, req.Page, req.PageSize)
	if err != nil && err != gorm.ErrRecordNotFound {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Data = response.AuditSearch{
		Total: count,
		List:  list,
	}
	return
}
func Create(ctx *gin.Context) (res controller.STDResponse, err error) {
	req := audit.Create{}
	//bind args
	if err = BindMultiForm(ctx, &req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	err = service.Create(req.AppID, req.AppType, req.FormData)
	if utils.IsMysqlDupKeyErr(err) {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = "唯一键冲突"
		return
	}
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
	}
	return
}

//O端审核
func UpdateState(ctx *gin.Context) (res controller.STDResponse, err error) {
	// 1.更新审核表
	// 2.更新enterprise表
	auditID, ok := ctx.Params.Get("audit_id")
	if !ok {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = "invalid params. No app_id provided"
		return
	}
	req := audit.UpdateState{}
	if err = BindJSON(ctx, &req); err != nil {
		return
	}
	rowCount, err := service.UpdateState(auditID, req.AppID, req.State, req.Comment)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Data = map[string]int64{
		"affected_rows": rowCount,
	}
	return
}
