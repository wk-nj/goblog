package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/password"
	"goblog/pkg/types"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(idStr string)  (User, error){
	var user User
	id := types.StringToUint64(idStr)
	err := model.DB.First(&user, id).Error
	if err != nil {
		logger.LogError(err)
		return user, err
	}
	return user, nil
}

func GetByEmail(email string)  (User, error){
	var user User
	err := model.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		logger.LogError(err)
		return user, err
	}
	return user, nil
}

func (user *User) ComparePassword(passwordRequest string)  bool{
	return password.CheckHash(passwordRequest, user.Password)
}

// All 获取所有用户数据
func All() ([]User, error) {
	var users []User
	if err := model.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}