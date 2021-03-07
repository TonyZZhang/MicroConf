package main

import (
	"MicroConf/ConfigService/config"
	"MicroConf/ConfigService/dao"
	"MicroConf/ConfigService/gee"
	"MicroConf/ConfigService/handler"
	"fmt"
	"log"
)

var conf config.Config

func main() {
	fmt.Println("start")
	conf.DB.URL = ""
	conf.DB.DriverName = "mysql"
	err := dao.Init(&conf)
	if err != nil {
		log.Fatal("DB connect error", err)
	}
	service := Router()
	service.Run(":9999")
}

func Router() *gee.Engine {
	r := gee.New()
	r.GET("/api/v1/get", handler.CreateApp)
	return r
}
