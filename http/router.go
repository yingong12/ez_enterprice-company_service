package http

import (
	"company_service/http/controller"
	"company_service/http/controller/audit"
	"company_service/http/controller/enterprise"
	"company_service/http/controller/value"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//routes
	router.POST("healthy", controller.Healthy)
	//swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
	// 企业模块
	groupEnterprise := router.Group("/enterprise")
	{
		groupEnterprise.POST("/search", enterprise.Search)        //企业信息搜索接口
		groupEnterprise.GET("/by_app_ids", enterprise.QueryByIDs) //根据企业id拿信息,批量
		groupEnterprise.POST("", enterprise.Create)               //新建企业 用于O端(zy要求)
		groupEnterprise.PUT(":app_id", enterprise.Update)         //更新企业信息
		//TODO:单独写状态接口因为查询状态较为频繁，减少网络请求数据量. 初期可以不使用
		// enterprise.GET("/state/:en_id") //获取企业状态
		// enterprise.PUT("/state/:en_id") //更新企业状态
	}
	// //机构模块
	// group := router.Group("/group")
	// {
	// 	// group.GET("") //获取机构信息
	// 	// group.GET("/enterprise/", controller.GetAssets) //获取机构拥有的企业id
	// }
	//审核模块
	groupAudit := router.Group("audits")
	{
		groupAudit.POST("", audit.Create)                     //提交审核 （涉及图片上传）
		groupAudit.GET("", audit.Search)                      //搜索审核,分页
		groupAudit.PUT("/:audit_id/state", audit.UpdateState) //审核通过，打回. 同步修改企业状态
	}
	// //评估模块
	groupValuate := router.Group("valuate")
	{
		groupValuate.POST("", value.Create) //提交估值
		groupValuate.GET("", value.Search)  //获取估值结果
		// groupValuate.GET(":id", value.Search)     //根据id获取估值结果
		// groupValuate.POST("export", value.Export) //导出 同步异步？
	}

	return
}
