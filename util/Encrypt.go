package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

// AES ECB 加密
func AESECBEncrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)

	plaintextBytes = PKCS5Padding(plaintextBytes, block.BlockSize())

	ciphertext := make([]byte, len(plaintextBytes))

	for bs, be := 0, block.BlockSize(); bs < len(plaintextBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(ciphertext[bs:be], plaintextBytes[bs:be])
	}

	// 返回 Base64 编码的密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AES ECB 解密
func AESECBDecrypt(key []byte, ciphertext string) (string, error) {
	// 解码 Base64 密文
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	plaintextBytes := make([]byte, len(ciphertextBytes))

	for bs, be := 0, block.BlockSize(); bs < len(ciphertextBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(plaintextBytes[bs:be], ciphertextBytes[bs:be])
	}

	plaintextBytes = PKCS5Unpadding(plaintextBytes)

	// 返回解密后的明文
	return string(plaintextBytes), nil
}

// PKCS5 填充
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS5 去除填充
func PKCS5Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
