package service

import (
	"company_service/http/request"
	"company_service/model"
	"company_service/repository"
)

func Search(rangeFilters []request.RangeFilter, textFilters []request.TextFilter, sort []request.Sort, page, pageSize int) (res []model.Enterprise, total int, err error) {
	res, err = repository.Search(rangeFilters, textFilters, sort, page, pageSize)
	if err != nil {
		return
	}
	// total, err = repository.Total(rangeFilters, textFilters, sort)
	return
}
