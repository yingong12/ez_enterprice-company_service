package main

import (
	"company_service/http"
	"company_service/library"
	"company_service/library/env"
	"company_service/logger"
	"company_service/providers"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

//加载硬编码字典
func loadStatics(fpaths []string, names []string) (err error) {
	for k, path := range fpaths {
		//相对路径
		if path[0] != '/' {
			wd, errWD := os.Getwd()
			if errWD != nil {
				return
			}
			path = wd + "/" + path
		}
		//绝对路径
		fd, errF := os.Open(path)
		if errF != nil {
			err = errF
			return
		}
		//如果是文件夹直接报错
		if fs, fsErr := fd.Stat(); fsErr != nil || fs.IsDir() {
			err = errors.New("open" + path + " failed")
			return
		}
		//decode文件内容
		switch names[k] {
		case "industry":
			err = json.NewDecoder(fd).Decode(&providers.IndustryDict)
			if err != nil {
				return err
			}
		case "district":
			err = json.NewDecoder(fd).Decode(&providers.DisrictDict)
			if err != nil {
				return err
			}
		default:
		}
		fd.Close()
	}
	return
}

//bootstrap providers,以及routines
func bootStrap() (err error) {
	//加载环境变量

	filePath := ".env"
	serverPort := 8687
	flag.StringVar(&filePath, "c", ".env", "配置文件")
	flag.IntVar(&serverPort, "p", 8687, "http端口")
	flag.Parse()
	if err = godotenv.Load(filePath); err != nil {
		return
	}
	log.Println("env loadded from file ", filePath)

	err, shutdownLogger := logger.Start()
	if err != nil {
		return
	}
	//加载static资源
	staticPath := env.GetStringVal("STATIC_PATH")
	if err = loadStatics([]string{staticPath + "/industry.json", staticPath + "/district.json"}, []string{"industry", "district"}); err != nil {
		return
	}
	//加载Redis连接池
	port := env.GetIntVal("REDIS_PORT_ENTERPRISE")
	poolSize := env.GetIntVal("REDIS_POOL_SIZE")
	redisConf := library.RedisConfig{
		ConnectionName: os.Getenv("SERVICE_NAME"),
		Addr:           os.Getenv("REDIS_ADDR_ENTERPRISE"),
		Port:           port,
		Password:       env.GetStringVal("REDIS_PSWD_ENTERPRISE"),
		DB:             0,
		PoolSize:       poolSize,
	}
	providers.RedisClient, err = library.NewRedisClient(&redisConf)
	if err != nil {
		return
	}
	// DB GORM初始化
	GormConfigs := []*library.GormConfig{
		{
			Receiver:       &providers.DBenterprise,
			ConnectionName: "gorm-core",
			DBName:         env.GetStringVal("DB_ENTERPRISE_RW_NAME"),
			Host:           env.GetStringVal("DB_ENTERPRISE_RW_HOST"),
			Port:           env.GetStringVal("DB_ENTERPRISE_RW_PORT"),
			UserName:       env.GetStringVal("DB_ENTERPRISE_RW_USERNAME"),
			Password:       env.GetStringVal("DB_ENTERPRISE_RW_PASSWORD"),
			MaxLifeTime:    env.GetIntVal("DB_MAX_LIFE_TIME"),
			MaxOpenConn:    env.GetIntVal("DB_MAX_OPEN_CONN"),
			MaxIdleConn:    env.GetIntVal("DB_MAX_IDLE_CONN"),
		},
	}

	for _, cfg := range GormConfigs {
		if cfg.Receiver == nil {
			return fmt.Errorf("[%s] config receiver cannot be nil", cfg.ConnectionName)
		}
		if *cfg.Receiver, err = library.NewGormDB(cfg); err != nil {
			return err
		}
		_, e := (*cfg.Receiver).DB.DB()
		if e != nil {
			return e
		}
	}
	//kafka
	pdcrs := []*library.ProducerConfig{
		{
			Brokers:  env.GetStringVal("KAFKA_BROKERS_VALUATE"),
			Receiver: &providers.ValProducer,
		},
	}
	for _, cfg := range pdcrs {
		if cfg.Receiver == nil {
			return fmt.Errorf("config receiver cannot be nil")
		}
		*cfg.Receiver, err = library.NewSyncProducer(cfg)
		if err != nil {
			return
		}
	}
	//http client
	httpClients := []*library.HttpClientConfig{
		{
			Name:     "static_server",
			BaseURL:  `http://` + env.GetStringVal("LB_STATIC_SERVICE") + "/",
			Receiver: &providers.HttpClientStatic,
		}, {
			Name:     "account_server",
			BaseURL:  `http://` + env.GetStringVal("LB_ACCOUNT_SERVICE") + "/",
			Receiver: &providers.HttpClientAccount,
		},
	}
	for _, cfg := range httpClients {
		if cfg.Receiver == nil {
			return fmt.Errorf("config receiver cannot be nil")
		}
		*cfg.Receiver = library.NewHttpClient(cfg)
		(**cfg.Receiver).BaseURL = cfg.BaseURL
	}
	//http server
	err, shutdownHttpServer := http.Start(serverPort)
	if err != nil {
		return
	}
	log.Println("Httpserver started ")

	//wait for sys signals
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-exitChan:
		log.Println("Doing cleaning works before shutdown...")
		shutdownLogger()
		shutdownHttpServer()
		log.Println("You abandoned me, bye bye", sig)
	}
	return
}
