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
	//写db
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

func UpdateState(auditID, appID string, state int) (rowCount int64, err error) {
	data := model.Enterprise{}
	data.State = int8(state)
	tx := providers.DBenterprise.Begin()
	defer func() {
		//rollback
		if err != nil || rowCount <= 0 {
			tx.Rollback()
		}
		tx.Commit()
	}()
	/*
		此处带入appid是为了校验该appid是否是库里存的appid
	*/
	//更新audit表
	rowCount, err = repository.UpdateState(auditID, appID, state)
	if rowCount <= 0 || err != nil {
		return
	}
	//更新enterprise表
	where := map[string]interface{}{"app_id": appID}
	rowCount, err = enterprise.SuperUpdate(where, data)
	return
}
