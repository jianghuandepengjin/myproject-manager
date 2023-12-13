package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	common "test.com/project-common"
	"test.com/project-common/logs"
	"test.com/project-user/internal/dao"
	"test.com/project-user/pkg/repo"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Redis,
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	//接受参数
	mobile := ctx.PostForm("mobile")
	//校验参数
	result := common.VerifyMobile(mobile)
	if !result {
		resp := common.Result{}
		ctx.JSON(http.StatusOK, resp.Fail(2001, "手机号不匹配"))
		//直接结束程序
		return
	}
	//生成验证码
	code := "123456"
	//调用第三方平台，短信服务发送验证码到手机上
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		//调用这个框架的
		zap.L().Debug("success send Debug log")
		//自己初始话的日志
		logs.LG.Info("success send InFo log")
		zap.L().Error("success send error log")
		//发送成功 存入redis
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码 restore is fail")
		}
	}()
	//把验证码存到缓存中
	ctx.JSON(200, "success this getCaptcha")
}
