package audit

import (
	"company_service/http/buz_code"
	"company_service/http/request/audit"
	"company_service/http/response"
	"company_service/logger"
	service "company_service/service/audit"
	"company_service/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func Search(ctx *gin.Context) {
	req := audit.Search{}
	if err := BindQuery(ctx, &req); err != nil {
		return
	}
	stateArr := []int{}
	if req.States != "" {
		//bind states
		stateStrArr := strings.Split(req.States, ",")
		for _, state := range stateStrArr {
			stateInt, err := strconv.Atoi(state)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": buz_code.CODE_INVALID_ARGS,
					"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
				})
				return
			}
			stateArr = append(stateArr, stateInt)
		}
	}

	list, count, err := service.Search(req.AppName, req.RegistrationNumber, req.AppID, stateArr, req.Page, req.PageSize)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": response.AuditSearch{
		Total: count,
		List:  list,
	}})
	return
}
func Create(ctx *gin.Context) {
	req := audit.Create{}
	//bind args
	if err := BindJSON(ctx, &req); err != nil {
		return
	}
	err := service.Create(req.AppID, req.AppType, req.FormData)
	if utils.IsMysqlDupKeyErr(err) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "唯一键冲突 " + err.Error(),
		})
		return
	}
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": buz_code.CODE_OK, "msg": "ok", "data": ""})
}

//O端审核
func UpdateState(ctx *gin.Context) {
	// 1.更新审核表
	// 2.更新enterprise表
	auditID, ok := ctx.Params.Get("audit_id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params. No app_id provided"),
		})
		return
	}
	req := audit.UpdateState{}
	if err := BindJSON(ctx, &req); err != nil {
		return
	}
	rowCount, err := service.UpdateState(auditID, req.AppID, req.State)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": buz_code.CODE_OK,
		"msg":  "ok",
		"data": map[string]int64{
			"affected_rows": rowCount,
		},
	})
	return
}
