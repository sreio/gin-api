package main

import (
	"fmt"
	"gin-api/models"
	"gin-api/pkg/logging"
	"gin-api/pkg/setting"
	"gin-api/routers"
	"log"
	"syscall"

	"github.com/fvbock/endless"
)

func main() {
	//初始化配置项
	setting.Setup()
	models.Setup()
	logging.Setup()

	
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// router := routers.InitRouter()
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}