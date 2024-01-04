package main

import (
	"github.com/gin-gonic/gin"
	_ "test.com/project-api/api"
	"test.com/project-api/config"
	"test.com/project-api/router"
	common "test.com/project-common"
)

func main() {
	r := gin.Default()
	//路由操作
	router.InitRoute(r)
	//优雅地启停操作
	common.Run(r, config.C.SC.Name, config.C.SC.Adds, nil)
}
