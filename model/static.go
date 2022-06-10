package model

// 硬编码城市经纬度信息
type District struct {
	Label    string      `json:"label"`
	Level    string      `json:"level"`
	Code     string      `json:"code"`
	Children []*District `json:"children"`
}

type IndustryDict struct {
	Text     string          `json:"text"`
	Value    string          `json:"value"`
	Children []*IndustryDict `json:"children"`
}
