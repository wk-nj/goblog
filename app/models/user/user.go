package user

import (
	"goblog/app/models"
)

type User struct {
	models.BaseModel
	Name     string `gorm:"column:name;type:varchar(100);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(100);default:NULL;unique;"`
	Password string `gorm:"column:password;type:varchar(255)"`
}
