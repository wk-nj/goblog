package user

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type User struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(100);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(100);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// Link 方法用来生成用户链接
func (user User) Link() string {
	return route.Name2URL("users.show", "id", user.GetStringID())
}