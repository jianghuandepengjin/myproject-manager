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

	////从配置中读取日志配置，初始化日志
	//lc := &logs.LogConfig{
	//	DebugFileName: "D:\\GO_WorkSpace\\src\\ms_project\\logs\\debug\\project-debug.log",
	//	InfoFileName:  "D:\\GO_WorkSpace\\src\\ms_project\\logs\\info\\project-info.log",
	//	WarnFileName:  "D:\\GO_WorkSpace\\src\\ms_project\\logs\\error\\project-error.log",
	//	MaxSize:       500,
	//	MaxAge:        28,
	//	MaxBackups:    3,
	//}
	//err := logs.InitLogger(lc)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//路由操作
	router.InitRoute(r)
	//gprc服务注册
	gc := router.ResgiterFunc()
	stop := func() {
		gc.Stop()
	}
	//优雅的启停操作
	common.Run(r, config.C.SC.Name, config.C.SC.Adds, stop)
}
