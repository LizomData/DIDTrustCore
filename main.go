package main

import (
	"DIDTrustCore/controller/didController"
	"DIDTrustCore/controller/fileUploadController"
	"DIDTrustCore/controller/sbomController"
	"DIDTrustCore/controller/userController"
	"DIDTrustCore/did_connnection"
	_ "DIDTrustCore/docs"
	"DIDTrustCore/routers"
	"DIDTrustCore/service"
	"DIDTrustCore/util/swagger"
	"fmt"
)

func main() {
	did_connnection.FabricClient = did_connnection.Connection()
	// 加载多个APP的路由配置
	routers.Include(userController.Routers, service.Routers, swagger.Routers, didController.Routers, sbomController.Routers, fileUploadController.Routers)
	// 初始化路由
	r := routers.Init()
	//初始化fabric客户端
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
