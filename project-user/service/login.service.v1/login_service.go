package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-user/internal/dao"
	"test.com/project-user/internal/data"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache        repo.Cache
	member       repo.MemberDao
	organization repo.Organization
}

// 这个类似service 层 new dao对象过来使用
func New() *LoginService {
	return &LoginService{
		cache:        dao.Redis,
		member:       dao.NewMeberDao(),
		organization: dao.NewOrganizationDao(),
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//校验参数
	//todo-为什么这里指针类型可以这样取值----值得思考
	mobile := msg.Mobile
	result := common.VerifyMobile(mobile)
	if !result {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	//生成验证码
	code := "123456"
	//调用第三方平台，短信服务发送验证码到手机上
	go func() {
		//time.Sleep(2 * time.Second)
		//log.Println("调用短信平台发送短信")
		////调用这个框架的
		//zap.L().Debug("success send Debug log")
		////自己初始话的日志
		//logs.LG.Info("success send InFo log")
		//zap.L().Error("success send error log")
		//发送成功 存入redis
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		//todo 先要判断redis中是否有数据，没数据菜存进去---其实这里也需要判断需求。 看是否需要判断---------如果超过某一段时间要重写发-就需要重写写。
		// 自己多根据实际日常生活去思考 写的逻辑是否可以优化
		// 还是得后端来，因为前端是限制不了的，人家重置请求就行，  不是按钮不让按 就不会再进来的
		err := ls.cache.Put(c, model.Register_key+msg.Mobile, code, 15*time.Minute)
		//todo 修改错误返回
		if err != nil {
			log.Println("验证码 restore is fail")
		}
	}()
	//把验证码存到缓存中
	return &CaptchaResponse{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *RegisterMessage) (*RegisterResponse, error) {
	c := context.Background()
	//todo test需要，先注釋了
	//判断数据是否都存在
	//判断验证码是否正确-从redis里面取
	//captche, err := ls.cache.Get(c, model.Register_key+msg.Mobile)
	//if err != nil {
	//	//todo 这里要分不同的 redis 错误
	//	return nil, errs.GrpcError(model.RedisError)
	//}
	//if captche != msg.Captcha {
	//	return nil, errs.GrpcError(model.CapchaError)
	//}

	//业务逻辑的校验（邮箱是哦福被注册，账号是否被注册， 手机号是否被注册）
	//todo 可优化放到一个函数
	conn := ls.member
	exist, err := conn.GetEmailFromMember(c, msg.Email)
	if err != nil {
		zap.L().Error("Db connect is fial", zap.Error(err))
		return nil, errs.GrpcError(model.DbError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailOfExistError)
	}
	exist, err = conn.GetPhoneFromMember(c, msg.Mobile)
	if err != nil {
		zap.L().Error("Db connect is fial", zap.Error(err))
		return nil, errs.GrpcError(model.DbError)
	}
	if exist {
		return nil, errs.GrpcError(model.PhoneOfExistError)
	}
	exist, err = conn.GetAccountFromMember(c, msg.Name)
	if err != nil {
		zap.L().Error("Db connect is fial", zap.Error(err))
		return nil, errs.GrpcError(model.DbError)
	}
	if exist {
		return nil, errs.GrpcError(model.AccountOfExistError)
	}

	//然后把数据都存到数据库中
	userInfo := data.Member{
		Email:    msg.Email,
		Name:     msg.Name,
		Password: msg.Password,
		Mobile:   msg.Mobile,
		Status:   1,
	}
	//copier.Copy(userInfo, msg)  ---todo 爲什麽問不能成功copy
	//1、先插入到用户表中
	_, err = conn.InsertUserTOMember(c, userInfo)
	if err != nil {
		return nil, errs.GrpcError(model.InsertOfUserError)
	}

	//2、然后插入到组织表中
	organizationConn := ls.organization
	organization := data.Organization{
		Id:         userInfo.Id,
		Name:       msg.Name + "各人主族",
		MemberId:   1,
		CreateTime: time.Now().UnixMilli(),
		Personal:   1,
	}
	_, err = organizationConn.InsertOrganization(c, organization)
	if err != nil {
		return nil, errs.GrpcError(model.InsertOfOrganizationError)
	}

	//直接返回成功
	return &RegisterResponse{}, nil
}
