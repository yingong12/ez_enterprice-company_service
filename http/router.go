package http

import (
	_ "company_service/docs" //引入swagger
	"company_service/http/controller"
	"company_service/http/controller/audit"
	"company_service/http/controller/enterprise"
	"company_service/http/controller/group"
	"company_service/http/controller/value"
	"company_service/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ControllerErrorLogger())
	//routes
	router.POST("health", controller.Health)
	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // register swagger
	// 企业模块
	groupEnterprise := router.Group("/enterprise")
	{
		groupEnterprise.POST("/search", controller.STDwrapperJSON(enterprise.Search))                         //企业信息搜索接口
		groupEnterprise.GET("/by_app_ids", controller.STDwrapperJSON(enterprise.QueryByIDs))                  //根据企业id拿信息,批量
		groupEnterprise.POST("", controller.STDwrapperJSON(enterprise.Create))                                //新建企业 用于O端(zy要求)
		groupEnterprise.PUT(":app_id", controller.STDwrapperJSON(enterprise.Update))                          //更新企业信息
		groupEnterprise.GET("get_industry_children", controller.STDwrapperJSON(enterprise.GetIndustryByCode)) //        根据industry id查询行业节点以及他的所有儿子
		groupEnterprise.GET("get_district_children", controller.STDwrapperJSON(enterprise.GetDistrictByCode)) //            根据地区代码查询地区节点以及他的所有儿子
	}
	// //机构模块
	groupGroup := router.Group("/group")
	{
		groupGroup.POST("", controller.STDwrapperJSON(group.Create))                            //新增机构
		groupGroup.GET("", controller.STDwrapperJSON(group.Search))                             //机构列表
		groupGroup.GET("get_children_multi", controller.STDwrapperJSON(group.GetChildrenMulti)) //获取机构关联的企业
		groupGroup.PUT(":id", controller.STDwrapperJSON(group.Update))                          //修改机构信息
	}
	//审核模块
	groupAudit := router.Group("audit")
	{
		groupAudit.POST("", controller.STDwrapperJSON(audit.Create))                     //提交审核 （涉及图片上传）
		groupAudit.GET("", controller.STDwrapperJSON(audit.Search))                      //搜索审核,分页
		groupAudit.PUT("/:audit_id/state", controller.STDwrapperJSON(audit.UpdateState)) //审核通过，打回. 同步修改企业状态
	}
	// //评估模块
	groupValuate := router.Group("valuate")
	{
		groupValuate.POST("", controller.STDwrapperJSON(value.Create)) //提交估值
		groupValuate.GET("", controller.STDwrapperJSON(value.Search))  //获取估值结果
		// groupValuate.GET(":id", value.Search)     //根据id获取估值结果
		// groupValuate.POST("export", value.Export) //导出 同步异步？
	}

	return
}
