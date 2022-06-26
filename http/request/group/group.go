package group

import "company_service/model"

type Search struct {
	AppIDs   string `form:"app_ids" exmple:"app1,app2,appp3"`
	Name     string `form:"name" exmple:"机构1"`
	Sort     uint8  `form:"sort" exmple:"排序方式,按拥有企业数量排序 0主键逆序 1-企业数量升序 2-企业数量降序"`
	Page     int    `form:"page" exmple:"1"`
	PageSize int    `form:"page_size" exmple:"10"`
}

type GetChildrenMulti struct {
	AppIDs string `json:"app_ids" exmple:"app1,app2,appp3"`
}

//
type Create struct {
	UID  string              `json:"uid"` /*用户id*/
	Data model.GroupMuttable //字段
}

type Update struct {
	Data model.GroupMuttable
}
