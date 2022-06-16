package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
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
	password := user.Password
	if password == passwordRequest {
		return true
	}
	return false
}