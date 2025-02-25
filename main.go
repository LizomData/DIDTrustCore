package main

import (
	"DIDTrustCore/controller/userController"
	_ "DIDTrustCore/docs"
	"DIDTrustCore/routers"
	"DIDTrustCore/service"
	"fmt"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {

	// 加载多个APP的路由配置
	routers.Include(userController.Routers, service.Routers)

	// 初始化路由
	r := routers.Init()

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
