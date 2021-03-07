package main

import (
	"AdminConf/config"
	"AdminConf/dao"
	"AdminConf/gee"
	"AdminConf/handler"
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
	r.POST("/api/v1/add", handler.CreateApp)
	return r
}
