package request

import "company_service/model"

type Search struct {
	Sort         []Sort        `json:"sort"`          //排序
	TextFilters  []TextFilter  `json:"textFilters"`   //全文搜索
	RangeFilters []RangeFilter `json:"range_filters"` //范围搜索
	Page         int           `json:"page"`          //页码
	PageSize     int           `json:"page_size"`     //分页大小
}

type RangeFilter struct {
	Type int `json:"type"` //0-注册资本 1-估值
	Gte  int `json:"gte"`  // >=
	Lte  int `json:"lte"`  // <=
}

type TextFilter struct {
	Type   int      `json:"type"` //0-行业代码
	Values []string `json:"values"`
}

type Sort struct {
	Type   uint8 `json:"type"`   //0-asc 1-desc
	Column int   `json:"column"` //0-注册资本 1-估值结果 2-名字
}

//
type Create struct {
	UID      string                   `json:"uid"`       //用户id
	ParentID string                   `json:"parent_id"` //机构id 非必填
	Data     model.EnterpriseMuttable //字段
}

type Update struct {
	Data model.EnterpriseMuttable
}
