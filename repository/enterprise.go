package repository

import (
	"company_service/http/request"
	"company_service/model"
	"company_service/providers"
	"company_service/utils"
)

func Search(rangeFilters []request.RangeFilter, textFilters []request.TextFilter, sort []request.Sort, page, pageSize int) (res []model.Enterprise, err error) {
	//转化filter 和 sort为sql
	en := []model.Enterprise{}
	tx := providers.DBenterprise.
		Debug().
		Table(model.GetEnterpriseTable())
	//全文搜索
	for _, v := range textFilters {
		p := utils.ParseFilter(v.Type)
		for _, v1 := range v.Values {
			tx.Where(p, v1)
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
	//Find
	tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&res)
	err = tx.Error
	res = en
	return
}
func Total(rangeFilters []request.RangeFilter, textFilters []request.RangeFilter, sort []request.Sort) (total int, err error) {
	return
}
