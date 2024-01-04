package router

import (
	"github.com/gin-gonic/gin"
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
	ro.Route(r)
}

var routers []Route

func InitRoute(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

// 传入多个值参数的方式  ...Route等价于 []Route
func Register(ro ...Route) {
	routers = append(routers, ro...)
}
