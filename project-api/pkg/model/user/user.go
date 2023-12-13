package user

import (
	"errors"
	common "test.com/project-common"
)

type RegisterReq struct {
	Email         string `json:"email" form:"email"`
	Name          string `json:"name" form:"name"`
	Password      string `json:"password" form:"password"`
	PasswordAgain string `json:"passwordagain" form:"passwordagain"`
	Mobile        string `json:"mobile" form:"mobile"`
	Captcha       string `json:"captcha" form:"captcha"`
}

// 这种小写只提供给内部调用，不给外部去用，外部提示就不会显示
func (r RegisterReq) verifyPassword() bool {
	return r.Password == r.PasswordAgain
}

func (r RegisterReq) Verify() error {
	if !common.VerifyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if !r.verifyPassword() {
		return errors.New("两次密码输入不一致")
	}
	return nil
}
