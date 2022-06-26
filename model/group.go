package model

import "time"

type Group struct {
	GroupMuttable
	GroupImmutable
	GroupBuzFields
}
type GroupBuzFields struct {
	LabelIndustry []string `json:"label_industry" gorm:"-"`
	LabelDistrict []string `json:"label_district" gorm:"-"`
}

//业务侧可create和update的字段
type GroupMuttable struct {
	Name                     string `gorm:"column:name" json:"name"`
	RegistrationNumber       string `gorm:"column:registration_number" json:"registration_number"`
	District                 string `gorm:"column:district" json:"district"`
	LegalRepresentative      string `gorm:"column:legal_representative" json:"legal_representative"`
	RegistrationAddress      string `gorm:"column:registration_address" json:"registration_address"`
	Industry                 string `gorm:"column:industry" json:"industry"`
	LicenseImg               string `gorm:"column:license_img" json:"license_img"`
	LegalRepresentativeIDImg string `gorm:"column:legal_representative_id_img" json:"legal_representative_id_img"`
	InVestScope              string `gorm:"column:invest_scope" json:"investment_areas"`
	Introduction             string `gorm:"column:introduction" json:"introduction"`
	ShareHolders             string `gorm:"column:share_holder_info" json:"share_holder_info"`
	RegisterCapital          int    `gorm:"column:register_capital" json:"register_capital"`
	CompanyType              int    `gorm:"column:company_type" json:"company_type"`
	Stage                    string `gorm:"column:invest_stage" json:"stage"`
	ShareHoldersJSON         string `gorm:"-" json:"shar_holders_json"`
	BusinessScope            string `gorm:"column:business_scope" json:"business_scope"`
	ChildrenCount            int    `gorm:"column:children_count" json:"children_count"`
}

//不允许业务update的字段
type GroupImmutable struct {
	AppID        string    `gorm:"column:app_id" json:"app_id"`
	UID          string    `gorm:"column:uid" json:"uid"`
	State        int8      `gorm:"column:state" json:"state"`                                     //审核状态
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 更新时间
	CreatedAtFmt string    `gorm:"-" json:"created_at"`                                           //返回给业务侧
	UpdatedAtFmt string    `gorm:"-" json:"udated_at"`
}

func GetGroupTable() string {
	return "t_group"
}
