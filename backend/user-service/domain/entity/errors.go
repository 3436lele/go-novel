package entity

import "errors"

var (
	ErrInvalidUsername = errors.New("用户名长度必须在3-50个字符之间")
	ErrInvalidEmail    = errors.New("邮箱不能为空")
	ErrInvalidPassword = errors.New("密码长度不能少于6个字符")
)
