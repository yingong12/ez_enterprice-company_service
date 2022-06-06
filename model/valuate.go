package model

import "time"

type Valuate struct {
	ValuateMuttable
	ValuateID    string    `gorm:"column:valuate_id" json:"valuate_id"`
	RequestedAt  time.Time `gorm:"column:requested_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"requested_at" `
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 更新时间
	CreatedAtFmt string    `gorm:"-" json:"created_at"`                                           //返回给业务侧
	UpdatedAtFmt string    `gorm:"-" json:"udated_at"`
}

//业务改变字段
type ValuateMuttable struct {
	AppID    string `gorm:"column:app_id" json:"app_id"`
	FormData string `gorm:"column:form_data" json:"form_data"`
	State    uint8  `gorm:"column:state" json:"state"`
}

//Kafka ValuateMessage
type ValuateMessage struct {
}

func GetValuateTable() string {
	return "t_valuates"
}
