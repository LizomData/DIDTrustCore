package userController

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util"
)

// 密钥（16 字节 = AES-128, 24 字节 = AES-192, 32 字节 = AES-256）
var Key_aes = []byte("d8d73c9f91ad4fc689cf1dac563dc906") // 16 字节密钥（AES-128）

func DecryptPassword(user *model.User) (string, error) {
	password_decrypt, err := util.AESECBDecrypt(Key_aes, user.Password)
	user.Password = password_decrypt
	return password_decrypt, err
}
