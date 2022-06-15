package repository

import (
	"company_service/http/request"
	"company_service/model"
	"company_service/providers"
	"company_service/utils"
	"fmt"

	"gorm.io/gorm"
)

//
func Search(tx *gorm.DB, rangeFilters []request.RangeFilter, textFilters []request.TextFilter, sort []request.Sort, page, pageSize int) (res []model.Enterprise, err error) {

	//TODO:需支持全文搜索
	for _, v := range textFilters {
		p := utils.ParseFilter(v.Type)
		for _, v1 := range v.Values {
			//这了用or连接
			q := fmt.Sprintf("%s like ?", p)
			tx.Where(q, "%"+v1+"%")
		}
	}
	//范围搜索
	for _, v := range rangeFilters {
		p := utils.ParseFilter(v.Type)
		if v.Gte >= 0 {
			tx.Where(p+" >= ? ", v.Gte)
		}
		if v.Gte <= 0 {
			tx.Where(p+" <= ? ", v.Lte)
		}
	}
	//排序
	if len(sort) > 0 {
		//排序
		orderClause := ""
		for _, v := range sort {
			orderClause += utils.ParseSortColumn(int(v.Type))
			if v.Type == 1 {
				orderClause += " DESC"
			}
			orderClause += " ,"
		}
		//去掉最后一个AND
		orderClause = orderClause[:len(orderClause)-1]
		tx.Order(orderClause)
	}
	//page>0启用分页
	if page > 0 {
		tx.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	tx.Find(&res)
	err = tx.Error
	return
}
func GetEnterpriseByAppIDs(appIDs []string) (res []model.Enterprise, err error) {
	tx := providers.DBenterprise.
		Table(model.GetEnterpriseTable())
	tx.Where("app_id IN ?", appIDs).Find(&res)
	err = tx.Error
	return
}
func GetEnterpriseByKey(key string, val interface{}) (res *model.Enterprise, err error) {
	res = &model.Enterprise{}
	tx := providers.DBenterprise.
		Table(model.GetEnterpriseTable())
	tx.Where(key, val).First(&res)
	err = tx.Error
	return
}
func Total(tx *gorm.DB, rangeFilters []request.RangeFilter, textFilters []request.TextFilter, sort []request.Sort) (total int64, err error) {
	for _, v := range textFilters {
		p := utils.ParseFilter(v.Type)
		for _, v1 := range v.Values {
			//这了用or连接
			q := fmt.Sprintf("%s like ?", p)
			tx.Where(q, "%"+v1+"%")
		}
	}
	//范围搜索
	for _, v := range rangeFilters {
		p := utils.ParseFilter(v.Type)
		if v.Gte >= 0 {
			tx.Where(p+" >= ? ", v.Gte)
		}
		if v.Gte <= 0 {
			tx.Where(p+" <= ? ", v.Lte)
		}
	}
	//排序
	if len(sort) > 0 {
		//排序
		orderClause := ""
		for _, v := range sort {
			orderClause += utils.ParseSortColumn(int(v.Type))
			if v.Type == 1 {
				orderClause += " DESC"
			}
			orderClause += " ,"
		}
		//去掉最后一个AND
		orderClause = orderClause[:len(orderClause)-1]
		tx.Order(orderClause)
	}
	tx.Count(&total)
	err = tx.Error
	return
}

func Create(appID, uid, pid string, data model.EnterpriseMuttable) (err error) {
	en := createEnterPriseEntityByMuttables(data)
	en.AppID = appID
	en.UID = uid
	en.ParentID = pid
	tx := providers.DBenterprise.
		Table(model.GetEnterpriseTable())
	tx.Create(en)
	err = tx.Error
	return
}

func createEnterPriseEntityByMuttables(data model.EnterpriseMuttable) (res model.Enterprise) {
	res.EnterpriseMuttable = data
	return
}

func Update(where map[string]interface{}, data model.EnterpriseMuttable) (affectedRows int64, err error) {
	en := createEnterPriseEntityByMuttables(data)
	tx := providers.DBenterprise.
		Table(model.GetEnterpriseTable())
	for k, v := range where {
		tx.Where(k, v)
	}
	tx.Updates(en)
	affectedRows = tx.RowsAffected
	err = tx.Error
	return
}

func SuperUpdate(where map[string]interface{}, data model.Enterprise) (affectedRows int64, err error) {
	tx := providers.DBenterprise.
		Table(model.GetEnterpriseTable())
	for k, v := range where {
		tx.Where(k, v)
	}
	tx.Updates(data)
	affectedRows = tx.RowsAffected
	err = tx.Error
	return
}
