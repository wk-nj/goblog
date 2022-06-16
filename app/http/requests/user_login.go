package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
	"net/url"
)

func ValidateLoginForm(data user.User)  url.Values{

	// 2. 表单规则
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
	}

	// 3. 定制错误消息
	messages := govalidator.MapData{
		"email": []string{
			"required:请填写登录账号",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请填写正确的邮箱地址",
		},
		"password": []string{
			"required:请填写密码",
			"min:长度需大于 6",
		},
	}

	// 4. 配置选项
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 标签标识符
		Messages: messages,
	}

	// 4. 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	return errs
}