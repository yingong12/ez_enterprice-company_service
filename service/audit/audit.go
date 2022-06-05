package audit

import (
	"company_service/model"
	"company_service/providers"
	enterprise "company_service/repository"
	repository "company_service/repository/audit"
	"company_service/utils"
)

func Create(appID string, appType uint8, formData string) (err error) {
	//获取唯一audit_id
	auditID := utils.GenerateAuditID()
	now := utils.GetNowFMT()
	return repository.Create(auditID, appID, appType, formData, now)
}

func Search(appName, registrationNumber, appID string, stateArr []int, page, pageSize int) (res []model.Audit, total int64, err error) {
	appIDs := []string{}
	res = make([]model.Audit, 0)
	//企业名称模糊查询
	if appName != "" {
		if appIDs, err = repository.GetAppIDsByNames(appName); err != nil {
			return
		}
	}
	//appid精确查询
	if appID != "" {
		appIDs = []string{appID}
	}
	//registration_number精确查询
	if registrationNumber != "" {
		appID, err = repository.GetAppIDsByRegistrationNumber(registrationNumber)
		if err != nil {
			return
		}
		appIDs = []string{appID}
	}
	// 拿total
	res, total, err = repository.Search(appIDs, stateArr, page, pageSize)
	return

}

func UpdateState(auditID string, state int) (rowCount int64, err error) {
	where := map[string]interface{}{"audit_id": auditID}
	data := model.Enterprise{}
	data.State = int8(state)
	tx := providers.DBenterprise.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	/*
		1.查auditID 找到appID.
		2.更新auditID state
		3.更新enterprise表
	*/
	//键入tx
	//更新enterprise表
	rowCount, err = enterprise.SuperUpdate(where, data)
	if err != nil {
		return
	}
	//更新audit表
	rowCount, err = repository.UpdateState(auditID, state)
	return
}
