package service

import "errors"

const (
	MinUserNameLength = 5
	MaxUserNameLength = 32
)

func CheckUser(userName string, password string) error {
	if userName == "" {
		return errors.New("用户名为空")
	}
	if len(userName) < MinUserNameLength {
		return errors.New("用户名小于最大长度")
	}
	if len(userName) > MaxUserNameLength {
		return errors.New("用户名大于最大长度")
	}
	if password == "" {
		return errors.New("密码为空")
	}
	if len(userName) < MinUserNameLength {
		return errors.New("密码小于最大长度")
	}
	if len(userName) > MaxUserNameLength {
		return errors.New("密码大于最大长度")
	}
	return nil
}
