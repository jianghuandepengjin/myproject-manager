package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Route interface {
	Route(r *gin.Engine)
}

type RegisterRoute struct {
}

func New() *RegisterRoute {
	return &RegisterRoute{}
}

func (*RegisterRoute) RouteImp(ro Route, r *gin.Engine) {
	//ro(r) 这里报错，这是一个接口传过来的对象，难怪下面觉得怪怪的。在这里面在调用方法
	ro.Route(r)
}

var routers []Route

// 首字母要大写
func InitRoute(r *gin.Engine) {
	//rg := New()
	//rg.RouteImp(&user.RouteUser{}, r)
	for _, ro := range routers {
		ro.Route(r)
	}
}

// 传入多个值参数的方式  ...Route等价于 []Route
func Register(ro ...Route) {
	routers = append(routers, ro...)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}
