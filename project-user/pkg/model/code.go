package model

import (
	"test.com/project-common/errs"
)

// 这里 var 和const 有什么区别 ---- 什么var还可以这样用
var (
	NoLegalMobile       = errs.NewError(2001, "手机号不合法") //手机号不合法
	RedisError          = errs.NewError(111001, "redis有误")
	CapchaError         = errs.NewError(2002, "captcha输入有误")
	DbError             = errs.NewError(111002, "数据库连接错误")
	EmailOfExistError   = errs.NewError(111003, "Email已经存在")
	PhoneOfExistError   = errs.NewError(111004, "Phone已经存在")
	AccountOfExistError = errs.NewError(111005, "Account已经存在")
)
