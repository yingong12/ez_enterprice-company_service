package model

// 硬编码城市经纬度信息
type District struct {
	Label    string      `json:"label,omitempty"`
	Level    string      `json:"level"`
	Code     string      `json:"code"`
	IsLeaf   bool        `json:"is_leaf"`
	Children []*District `json:"children,omitempty"`
}

type IndustryDict struct {
	Label    string          `json:"label,omitempty"`
	Code     string          `json:"code"`
	IsLeaf   bool            `json:"is_leaf"`
	Children []*IndustryDict `json:"children,omitempty"`
}
