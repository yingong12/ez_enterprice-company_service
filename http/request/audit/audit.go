package audit

type Create struct {
	AppID    string `form:"app_id"`
	AppType  uint8  `form:"app_type"`
	FormData string `form:"form_data"` //审核表单信息
}

type Search struct {
	States             string `json:"states" form:"states"` //状态
	Page               int    `json:"page" form:"page"`
	PageSize           int    `json:"page_size" form:"page_size"`
	AppID              string `json:"app_id" form:"app_id"`                           //企业id精确查询
	AppName            string `json:"app_name" form:"app_name"`                       //企业名字模糊查询
	RegistrationNumber string `json:"registration_number" form:"registration_number"` //注册号
	AuditIDs           string `json:"audit_ids" form:"audit_ids"`                     //审核id

}
type UpdateState struct {
	AppID   string `json:"app_id"`
	State   int    `json:"state"`   //
	Comment string `json:"comment"` //管理员备注
}
