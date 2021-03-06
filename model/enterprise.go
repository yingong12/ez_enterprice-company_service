package model

import "time"

type Enterprise struct {
	EnterpriseMuttable
	EnterpriseImmutable
	EnterpriseBuzFields
}
type EnterpriseBuzFields struct {
	LabelIndustry []string `json:"label_industry" gorm:"-"`
	LabelDistrict []string `json:"label_district" gorm:"-"`
}

//业务侧可create和update的字段
type EnterpriseMuttable struct {
	Name                     string `gorm:"column:name" json:"name"`
	RegistrationNumber       string `gorm:"column:registration_number" json:"registration_number"`
	District                 string `gorm:"column:district" json:"district"`
	LegalRepresentative      string `gorm:"column:legal_representative" json:"legal_representative"`
	RegistrationAddress      string `gorm:"column:registration_address" json:"registration_address"`
	Industry                 string `gorm:"column:industry" json:"industry"`
	LicenseImg               string `gorm:"column:license_img" json:"license_img"`
	LegalRepresentativeIDImg string `gorm:"column:legal_representative_id_img" json:"legal_representative_id_img"`
	BusinessScope            string `gorm:"column:business_scope" json:"business_scope"`
	Introduction             string `gorm:"column:introduction" json:"introduction"`
	ShareHolders             string `gorm:"column:share_holder_info" json:"share_holder_info"`
	RegisterCapital          int    `gorm:"column:register_capital" json:"register_capital"`
	CompanyType              int    `gorm:"column:company_type" json:"company_type"`
	EstimateValue            int    `gorm:"column:estimate_value" json:"estimate_value"`
	Stage                    int8   `gorm:"column:stage" json:"stage"`
	ShareHoldersJSON         string `gorm:"-" json:"shar_holders_json"`
	State                    int8   `gorm:"column:state" json:"state"` //审核状态
}

//不允许业务update的字段
type EnterpriseImmutable struct {
	AppID        string    `gorm:"column:app_id" json:"app_id"`
	UID          string    `gorm:"column:uid" json:"uid"`
	ParentID     string    `gorm:"column:parent_id" json:"parent_id"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 更新时间
	CreatedAtFmt string    `gorm:"-" json:"created_at"`                                           //返回给业务侧
	UpdatedAtFmt string    `gorm:"-" json:"udated_at"`
}

func GetEnterpriseTable() string {
	return "t_enterprise"
}

//
type Audit struct {
	AuditID      string    `gorm:"column:audit_id" json:"audit_id"`
	AppID        string    `gorm:"column:app_id" json:"app_id"`
	AppType      int8      `gorm:"column:app_type" json:"app_type"`
	FormData     string    `gorm:"column:form_data" json:"form_data"`
	State        uint8     `gorm:"column:state" json:"state"`
	RequestedAt  string    `gorm:"column:requested_at" json:"requested_at"`
	Comment      string    `gorm:"column:comment" json:"comment"`                                 //不通过理由等。
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 更新时间
	CreatedAtFmt string    `gorm:"-" json:"created_at"`                                           //返回给业务侧
	UpdatedAtFmt string    `gorm:"-" json:"udated_at"`
}

const ENTERPRISE_STATE_DELETED = -1

func GetAuditTable() string {
	return "t_audit_record"

}
