package main

import (
	"DIDTrustCore/controller/userController"
	_ "DIDTrustCore/docs"
	"DIDTrustCore/routers"
	"DIDTrustCore/service"
	"DIDTrustCore/util/swagger"
	"fmt"
)

func main() {

	// 加载多个APP的路由配置
	routers.Include(userController.Routers, service.Routers, swagger.Routers)

	// 初始化路由
	r := routers.Init()

	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
