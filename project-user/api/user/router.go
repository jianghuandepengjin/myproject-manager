package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-user/router"
)

// 只有在上面 import有引用这个文件，才会生效
func init() {
	log.Println("run this way")
	router.Register(&RouteUser{})
}

type RouteUser struct {
}

func (*RouteUser) Route(r *gin.Engine) {
	//在同一个包下面，所以不用user.   所以一会只看上面的package的值 不看目录结构
	h := New()
	//可以这样直接调用函数很奇怪--这是解决把函数作为参数
	r.POST("project/login/getCaptcha", h.getCaptcha)
}
