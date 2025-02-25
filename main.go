package main

import (
	"DIDTrustCore/controller/userController"
	"DIDTrustCore/routers"
	"DIDTrustCore/service"
	"fmt"
)

func main() {

	// 加载多个APP的路由配置
	routers.Include(userController.Routers, service.Routers)

	// 初始化路由
	r := routers.Init()

	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
