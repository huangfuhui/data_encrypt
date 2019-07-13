package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES加密
func aesEncrypt(data, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	data = PKCS5Padding(data, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	dataEncrypted := make([]byte, len(data))
	blockMode.CryptBlocks(dataEncrypted, data)

	// 使用base64编码加密结果，方便输出查看
	base64Encode := base64.StdEncoding.EncodeToString(dataEncrypted)

	return base64Encode, nil
}

// AES解密
func aesDecrypt(encrypted string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	base64Decode, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originData := make([]byte, len(base64Decode))
	blockMode.CryptBlocks(originData, base64Decode)
	originData = PKCS5UnPadding(originData)

	return string(originData), nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
