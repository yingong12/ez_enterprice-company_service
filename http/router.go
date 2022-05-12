package http

import (
	"account_service/http/controller"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//routes
	router.POST("healthy", controller.Healthy)
	//swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
	//登录模块
	auth := router.Group("/auth")
	{

		/**
		1. 登录态校验  token-> code 0:uid,app_id  code 1: 过期   code 2:wrong token code 3:token missing
		2. 注册： username,pswd,phone(带验证码) -> token
		3. 手机登录： phone#, veriCode -> token
		*/
		auth.GET("/check", controller.Check)                     //校验登录态
		auth.POST("/signin/username", controller.SignInUsername) //用户名登录
		auth.POST("/signin/sms", controller.SignInSMS)           //手机登录
		auth.POST("/signup/username", controller.SignUpUsername) //用户名注册
		auth.POST("/signup/sms", controller.SignUpSMS)           //手机注册
		// flow.Use(providers.ServiceOrder())
	}
	//账号模块
	account := router.Group("/account")
	{
		auth.GET("/assets", controller.GetAssets)               // 根据uid，查询拥有的企业，机构
		account.PUT("/username/:uid", controller.BindUsername)  //绑定用户名
		account.PUT("/phone/:uid", controller.BindPhone)        //绑定手机号
		account.PUT("/lock_unlock/:uid", controller.LockUnlock) //冻结，解冻用户
		auth.PUT("/pswd/:uid", controller.UpdatePswd)           //修改密码
		//TODO: 明天想重置密码怎么弄
		// auth.POST("/reset_pswd") //重置密码 需带一个special token
		//查询角色
	}
	//sms验证码模块
	sms := router.Group("sms")
	{
		sms.POST("/ask_code", controller.AskCode) //向sms服务申请验证码
	}
	return
}
