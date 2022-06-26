package providers

import (
	"company_service/library"
	"company_service/model"
)

var RedisClient *library.RedisClient

var DBenterprise *library.GormDB

//估值队列生产者
var ValProducer *library.SyncProducer

//static 服务
var HttpClientStatic *library.HttpClient
var HttpClientAccount *library.HttpClient

var IndustryDict model.IndustryDict
var DisrictDict model.District
