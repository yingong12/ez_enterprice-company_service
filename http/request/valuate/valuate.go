package valuate

import "company_service/model"

type Create struct {
	model.ValuateMuttable
}
type Search struct {
	AppID    string `form:"app_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
