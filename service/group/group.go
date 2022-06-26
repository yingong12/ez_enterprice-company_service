package group

import (
	"company_service/model"
	"company_service/providers"
	"company_service/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(uid string, data model.GroupMuttable) (err error) {
	log.Println(uid, data)
	appID := utils.GenerateGroupID()
	en := model.Group{}
	en.AppID = appID
	en.GroupMuttable = data
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	tx.Create(en)
	err = tx.Error
	return
}

func Update(ctx *gin.Context) (err error) {
	return
}

func Search(appIDs []string, name string, sort uint8, page int, pageSize int) (res []model.Group, total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	if len(appIDs) > 0 && appIDs[0] != "" {
		tx.Where("app_id in ?", appIDs)
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
	total, err = getTotal(appIDs, name, sort)
	return
}
func getTotal(appIDs []string, name string, sort uint8) (total int64, err error) {
	tx := providers.DBenterprise.Table(model.GetGroupTable())
	if len(appIDs) > 0 {
		tx.Where("app_id in ?", appIDs)
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
func ChilrenInfo() (err error) {
	return
}
