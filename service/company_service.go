package service

import (
	"company_service/http/buz_code"
	"company_service/http/request"
	"company_service/model"
	"company_service/providers"
	"company_service/repository"
	"company_service/utils"
	"errors"

	"gorm.io/gorm"
)

func Search(rangeFilters []request.RangeFilter, textFilters []request.TextFilter, sort []request.Sort, page, pageSize int) (res []model.Enterprise, total int, err error) {
	res, err = repository.Search(rangeFilters, textFilters, sort, page, pageSize)
	if err != nil {
		return
	}
	// total, err = repository.Total(rangeFilters, textFilters, sort)
	return
}

func Create(uid string, pid string, data model.EnterpriseMuttable) (buzCode buz_code.Code, msg string, err error) {
	//查是否已存在
	tx := providers.DBenterprise.Begin()

	defer func() {
		buzCode = buz_code.CODE_ENTERPRISE_CREATE_FAILED
		if err != nil {
			msg = err.Error()
			if msg == "该用户已经注册过公司" {
				buzCode = buz_code.CODE_ENTERPRISE_CREATE_FAILED
			}
			if utils.IsMysqlDupKeyErr(err) {
				buzCode = buz_code.CODE_ENTERPRISE_CREATE_FAILED
				msg = "唯一键冲突"
			}
			return
		}
		msg = "ok"
		tx.Commit()
		return
	}()
	//没被注册才继续
	_, err = repository.GetEnterpriseByKey("uid", uid)
	if err != gorm.ErrRecordNotFound {
		err = errors.New("该用户已经注册过公司")
		return
	}
	//能新建
	appID := utils.GenerateAppID()
	err = repository.Create(appID, uid, pid, data)
	return
}
func Update(appID string, data model.EnterpriseMuttable) (rows int64, err error) {
	//没被注册才继续
	where := map[string]interface{}{
		"app_id": appID,
	}
	rows, err = repository.Update(where, data)
	return
}

func GetByAppIDs(appIDs []string) (res []model.Enterprise, err error) {
	return repository.GetEnterpriseByAppIDs(appIDs)
}
