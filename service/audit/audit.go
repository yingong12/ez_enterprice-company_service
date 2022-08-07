package audit

import (
	"company_service/model"
	"company_service/providers"
	enterprise "company_service/repository"
	repository "company_service/repository/audit"
	group "company_service/service/group"
	"company_service/utils"
	"encoding/json"

	"gorm.io/gorm"
)

// 审核中
const STATE_AUDITING int8 = 0

func Create(appID string, appType int8, formData string) (err error) {
	//获取唯一audit_id
	auditID := utils.GenerateAuditID()
	now := utils.GetNowFMT()
	//写db
	where := map[string]interface{}{
		"app_id": appID,
	}
	//TODO: 这里用事务
	//更新企业，机构表
	if appType == 1 {
		//机构
		muData := model.GroupMuttable{}
		//解析表单
		if err = json.Unmarshal(([]byte)(formData), &muData); err != nil {
			return
		}
		_, err = group.Update(appID, muData, STATE_AUDITING)
	} else {

		muData := model.EnterpriseMuttable{}
		//解析表单
		if err = json.Unmarshal(([]byte)(formData), &muData); err != nil {
			return
		}
		_, err = enterprise.Update(where, muData, STATE_AUDITING)
	}
	if err != nil {
		return
	}
	//更新audit表
	repository.Create(auditID, appID, appType, formData, now)
	return
}

// Search 根据企业机构名称，注册表等信息查询appID 再去audit表里查审核。
func Search(appName, registrationNumber, appID string, stateArr []int, page, pageSize int) (res []model.Audit, total int64, err error) {
	appIDs := []string{}
	res = make([]model.Audit, 0)

	//appid精确查询
	if appID != "" {
		appIDs = []string{appID}
		goto final
	}
	//审核名称模糊查询
	if appName != "" {
		//查企业
		if appIDs, err = enterprise.GetAppIDsByNames(appName); err != nil {
			return
		}
		//查机构
		groupAppIDs := []string{}
		groupAppIDs, err = func(name string) (appIDs []string, err error) {
			res := []model.Enterprise{}
			tx := providers.DBenterprise.Table(model.GetGroupTable())
			tx.Select("app_id").Where("name like ?", "%"+name+"%").Find(&res)
			for _, v := range res {
				appIDs = append(appIDs, v.AppID)
			}
			err = tx.Error
			return
		}(appName)
		if err != nil {
			return
		}
		appIDs = append(appIDs, groupAppIDs...)
		// 条件之间and连接
		if len(appIDs) == 0 {
			return
		}
	}

	//registration_number精确查询
	if registrationNumber != "" {
		var en *model.Enterprise
		en, err = enterprise.GetEnterpriseByKey("registration_number", registrationNumber)
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		if err != gorm.ErrRecordNotFound {
			appIDs = []string{en.AppID}
			//因为注册号是唯一的，所以企业找到了就不去机构找了
			goto final
		}
		//查机构
		var gpEn *model.Group
		gpEn, err = func(registrationNumber string) (gpEn *model.Group, err error) {
			gpEn = &model.Group{}
			tx := providers.DBenterprise.Table(model.GetGroupTable())
			tx.Select("app_id").Where("registration_number", registrationNumber).First(&gpEn)
			err = tx.Error
			return
		}(registrationNumber)
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		if err != gorm.ErrRecordNotFound {
			appIDs = append(appIDs, gpEn.AppID)
		}
		if len(appIDs) == 0 {
			return
		}
	}
	// 拿total
final:
	res, total, err = repository.Search(appIDs, stateArr, page, pageSize)
	return

}

func UpdateState(auditID, appID string, appType uint16, state int, comment string) (rowCount int64, err error) {
	data := model.Enterprise{}
	data.State = int8(state)
	tx := providers.DBenterprise.Begin()
	defer func() {
		//rollback
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	/*
		此处带入appid是为了校验该appid是否是库里存的appid
	*/
	//更新audit表
	rowCount, err = repository.UpdateState(tx, auditID, appID, comment, state)
	if rowCount <= 0 || err != nil {
		return
	}
	where := map[string]interface{}{"app_id": appID}
	//机构
	if appType == 1 {
		rowCount, err = updateGroup(tx, where, data)
		return
	}
	//更新enterprise表
	rowCount, err = enterprise.SuperUpdate(tx, where, data)
	return
}

func updateGroup(tx *gorm.DB, where map[string]interface{}, data model.Enterprise) (affectedRows int64, err error) {
	tx = tx.Table(model.GetGroupTable())
	for k, v := range where {
		tx = tx.Where(k, v)
	}
	tx = tx.Updates(data)
	affectedRows = tx.RowsAffected
	err = tx.Error
	return
}
