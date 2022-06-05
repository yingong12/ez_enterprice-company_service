package response

import "company_service/model"

type Search struct {
	List  []model.Enterprise `json:"list"`
	Total int                `json:"total"`
}

type AuditSearch struct {
	List  []model.Audit `json:"list"`
	Total int64         `json:"total"`
}
