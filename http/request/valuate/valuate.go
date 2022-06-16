package valuate

import "company_service/model"

type Create struct {
	model.ValuateMuttable
}
type Search struct {
	AppID    string `form:"app_id"`    /*企业ID*/
	Page     int    `form:"page"`      /*页*/
	PageSize int    `form:"page_size"` /*页大小*/
}
