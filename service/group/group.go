package group

import (
	"company_service/model"
	"company_service/providers"
	"company_service/utils"
	"fmt"
	"log"
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

func Update(appID string, data model.GroupMuttable) (rf int64, err error) {
	log.Println(appID, data)
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	tx.Where("app_id", appID)
	tx.Updates(data)
	err = tx.Error
	rf = tx.RowsAffected
	return
}

func Search(appID string, name string, sort uint8, page int, pageSize int) (res []model.Group, total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	if appID != "" {
		tx.Where("app_id", appID)
	}
	if len(name) > 0 {
		tx.Where("name like ?", "%"+name+"%")
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
	tx.Order(sortClause)
	//page>0启用分页
	if page > 0 {
		tx.Offset((page - 1) * pageSize).Limit(pageSize)
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
	tx.Find(&res)
	err = tx.Error
	if err != nil {
		return
	}
	total, err = getChildrenTotal(appID)
	return
}
func getChildrenTotal(appID string) (total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetEnterpriseTable())
	tx.Where("parent_id", appID)
	tx.Count(&total)
	err = tx.Error
	return
}
