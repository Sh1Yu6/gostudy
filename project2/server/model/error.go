package model

import "errors"

var (
	ErrorUserNoExitsts = errors.New("用户不存在!")
	ErrorUserExitsts   = errors.New("用户已存在!")
	ErrorUserPwd       = errors.New("密码不正确!")
)
