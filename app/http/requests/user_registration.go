package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
	"net/url"
)

func ValidateRegistrationForm(data user.User)  url.Values{

	// 2. 表单规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:2,20", "not_exists:users,name"},
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	// 3. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
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

	// 5. 因 govalidator 不支持 password_confirm 验证，我们自己写一个
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}

	return errs
}