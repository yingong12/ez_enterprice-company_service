package model

type Enterprise struct {
	AppID                    string `gorm:"column:app_id"`
	UID                      string `gorm:"column:uid"`
	Name                     string `gorm:"column:name"`
	RegistrationNumber       string `gorm:"column:registration_number"`
	District                 string `gorm:"district"`
	LegalRepresentative      string `gorm:"legal_representative"`
	RegisterCapital          string `gorm:"register_capital"`
	RegisterationAddress     string `gorm:"registeration_address"`
	CompanyType              string `gorm:"company_type"`
	Industry                 string `gorm:"industry"`
	LegalRepresentativeIDImg string `gorm:"legal_representative_id_img"`
	BusinessScope            string `gorm:"business_scope"`
	Introduction             string `gorm:"introduction"`
	License_img              string `gorm:"license_img"`
	Estimate_value           string `gorm:"estimate_value"`
	Stage                    string `gorm:"stage"`
	ShareHolders             string `gorm:"share_holders"`
	ShareHolderProportion    string `gorm:"share_holder_proportion"`
	ParentID                 string `gorm:"parent_id"`
	CreatedAt                string `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt                string `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"` // 更新时间
}

func GetEnterpriseTable() string {
	return "t_enterprise"
}
