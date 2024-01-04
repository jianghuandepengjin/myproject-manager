package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

// 只有在上面 import有引用这个文件，才会生效
func init() {
	log.Println("run this way")
	router.Register(&RouteUser{})
}

type RouteUser struct {
}

func (*RouteUser) Route(r *gin.Engine) {
	//初始化api的具体逻辑
	h := New()

	//初始化grpc客户端
	initGrpcUserClient()

	//可以这样直接调用函数很奇怪--这是解决把函数作为参数
	r.POST("project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
}
