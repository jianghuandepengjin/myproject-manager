package main

import (
	"github.com/gin-gonic/gin"
	common "test.com/project-common"
	_ "test.com/project-user/api"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	//路由操作
	router.InitRoute(r)
	//grpc服务注册
	gc := router.ResgiterFunc()
	stop := func() {
		gc.Stop()
	}
	//优雅地启停操作
	common.Run(r, config.C.SC.Name, config.C.SC.Adds, stop)
}
