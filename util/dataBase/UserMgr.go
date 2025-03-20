package dataBase

import "DIDTrustCore/model"

// 创建用户
func CreateUser(user model.User) error {
	result := db.Create(&user)
	return result.Error
}

// 查询用户
func FindUser(username string) (bool, model.User) {
	var user model.User
	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return false, user
	}
	return true, user

}

// 查询用户
func FindUserById(userId uint) (bool, model.User) {
	var user model.User
	result := db.First(&user, "id = ?", userId)
	if result.Error != nil {
		return false, user
	}
	return true, user

}

// 更新用户
func UpdateUser(user model.User) error {
	result := db.Save(&user)
	return result.Error
}

// 删除用户
func DeleteUser(user model.User) error {
	result := db.Delete(&user)
	return result.Error

}
