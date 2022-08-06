package group

import (
	"company_service/model"
	"company_service/providers"
	"company_service/utils"
	"fmt"
)

func Create(appID, uid string, data model.GroupMuttable) (err error) {
	if appID == "" {
		appID = utils.GenerateGroupID()
	}
	en := model.Group{}
	en.AppID = appID
	en.GroupMuttable = data
	en.UID = uid
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	tx.Create(en)
	err = tx.Error
	return
}

func Update(appID string, data model.GroupMuttable, state int8) (rf int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	tx = tx.Where("app_id", appID).Updates(data)
	if state != -1 {
		tx = tx.Update("state", state)
	}
	err = tx.Error
	rf = tx.RowsAffected
	return
}

func Search(appID string, name string, sort uint8, page int, pageSize int) (res []model.Group, total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	if appID != "" {
		tx = tx.Where("app_id", appID)
	}
	if len(name) > 0 {
		tx = tx.Where("name like ?", "%"+name+"%")
	}
	sortClause := ""
	if sort == 0 {
		//默认主键逆序
		sortClause = fmt.Sprintf("id DESC")
	} else if sort == 1 {
		sortClause = fmt.Sprintf("children_count ASC")
	} else {
		sortClause = fmt.Sprintf("children_count DESC")
	}
	tx = tx.Where("state <> ?", model.ENTERPRISE_STATE_DELETED)
	tx = tx.Order(sortClause)
	//page>0启用分页
	if page > 0 {
		tx = tx.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	tx.Find(&res)
	err = tx.Error
	if err != nil {
		return
	}
	total, err = getTotal(appID, name, sort)
	return
}
func getTotal(appID, name string, sort uint8) (total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	if appID != "" {
		tx.Where("app_id", appID)
	}
	if len(name) > 0 {
		tx.Where("name like ?", "%"+name+"%")
	}
	tx = tx.Where("state <> ?", model.ENTERPRISE_STATE_DELETED)
	sortClause := ""
	if sort == 0 {
		//默认主键逆序
		sortClause = fmt.Sprintf("id DESC")
	} else if sort == 1 {
		sortClause = fmt.Sprintf("children_count ASC")
	} else {
		sortClause = fmt.Sprintf("children_count DESC")
	}
	tx.Order(sortClause)
	tx.Count(&total)
	err = tx.Error
	return
}

//
func ChilrenInfo(appID string, page, pageSize int) (res []model.Enterprise, total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetEnterpriseTable())
	tx.Where("parent_id", appID)
	//page>0启用分页
	if page > 0 {
		tx.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	tx = tx.Where("state <> ?", model.ENTERPRISE_STATE_DELETED)
	tx.Find(&res)
	err = tx.Error
	if err != nil {
		return
	}
	for k := range res {
		if m := utils.DFSDistrict(&providers.DisrictDict, res[k].District); m != nil {
			res[k].LabelDistrict = []string{m.Label}
		}
		if m := utils.DFSIndustry(&providers.IndustryDict, res[k].Industry); m != nil {
			res[k].LabelIndustry = []string{m.Label}
		}
	}
	total, err = getChildrenTotal(appID)
	return
}
func getChildrenTotal(appID string) (total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetEnterpriseTable())
	tx.Where("parent_id", appID)
	tx = tx.Where("state <> ?", model.ENTERPRISE_STATE_DELETED)
	tx.Count(&total)
	err = tx.Error
	return
}
