package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	login "test.com/project-api/api/user/user_grpc"
	"test.com/project-api/pkg/model/user"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"time"
	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

type User struct {
	id int
}

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	//todo 不能没一个api这里都要这样设置把,需要提取
	result := &common.Result{}

	//接受参数
	mobile := ctx.PostForm("mobile")
	if mobile == "" {
		//todo 把这些常数整理到一个文件里面
		ctx.JSON(http.StatusOK, result.Fail(401, "请传入有效参数"))
		return
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	captchaResponse, err := UserClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		//issue：为什么不放外面，要放里面---如果都没err 不都浪费空间
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(captchaResponse.Code))
}

func (h *HandlerUser) register(ctx *gin.Context) {
	result := &common.Result{}
	//接受参数
	var req user.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		//这种 return hi是值得学习的， 不使用else  --- 如果是循环的化也可以return， 还有break等等
		return
	}
	//校验参数
	err = req.Verify()
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		return
	}
	//调grpc服务，完成注册
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	resp, err := UserClient.Register(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//err = copier.Copy(msg, req)
	//返回注册结果
	ctx.JSON(http.StatusOK, result.Success(resp))
}
