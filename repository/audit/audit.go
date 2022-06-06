package audit

import (
	"company_service/model"
	"company_service/providers"
	"company_service/repository"
)

func Create(auditID, appID string, appType uint8, formData string, requestedAt string) (err error) {
	//插入, 删掉多余的字段
	tx := providers.DBenterprise.Table(model.GetAuditTable())
	en := model.Audit{
		AuditID:     auditID,
		AppID:       appID,
		AppType:     appType,
		FormData:    formData,
		RequestedAt: requestedAt,
	}
	tx.Create(en)
	err = tx.Error
	return
}

func Search(appIDs []string, stateArr []int, page, pageSize int) (res []model.Audit, count int64, err error) {
	tx := providers.DBenterprise.Table(model.GetAuditTable())
	res = make([]model.Audit, 0)
	//单个和多个分开写sql
	if len(appIDs) > 0 {
		if len(appIDs) == 1 {
			tx.Where("app_id = ?", appIDs[0])
		}
		if len(appIDs) > 1 {
			tx.Where("app_id IN ?", appIDs)
		}
	}
	if len(stateArr) > 0 {
		tx.Where("state in ?", stateArr)
	}
	tx.Count(&count)
	if count == 0 {
		return
	}
	tx.Order("requested_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize)
	tx.Find(&res)
	err = tx.Error
	return
}

func UpdateState(auditID, appID string, state int) (rowCount int64, err error) {
	tx := providers.DBenterprise.Table(model.GetAuditTable())
	tx.Where("audit_id", auditID)
	tx.Where("app_id", appID)
	tx.Update("state", state)
	rowCount = tx.RowsAffected
	err = tx.Error
	return
}

func GetAppIDsByNames(name string) (appIDs []string, err error) {
	if name == "" {
		return
	}
	res := []model.Enterprise{}
	tx := providers.DBenterprise.Table(model.GetEnterpriseTable())
	tx.Where("name like ?", "%"+name+"%").Find(&res)
	for _, v := range res {
		appIDs = append(appIDs, v.AppID)
	}
	err = tx.Error
	return
}

func GetAppIDsByRegistrationNumber(registrationNumber string) (appID string, err error) {
	if registrationNumber == "" {
		return
	}
	en, err := repository.GetEnterpriseByKey("registration_number", registrationNumber)
	if err != nil {
		return
	}
	appID = en.AppID
	return
}

func Total(appIDs []string) (total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetAuditTable())
	if len(appIDs) > 0 {
		tx.Where("app_id in ?", appIDs)
	}
	tx.Count(&total)
	err = tx.Error
	return
}
