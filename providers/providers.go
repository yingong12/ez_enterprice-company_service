package providers

import "company_service/library"

var RedisClient *library.RedisClient

var DBenterprise *library.GormDB

//估值队列生产者
var ValProducer *library.SyncProducer

//static 服务
var HttpClientStatic *library.HttpClient
