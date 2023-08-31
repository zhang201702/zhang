package zcrypto

import (
	"encoding/base64"
	"math/rand"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
)

/*
ASCII :33 到 126
*/
func Random(length int) []byte {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(rune(33 + rand.Intn(93)))
	}
	return result
}

func Encrypt(key, data string) []byte {
	iv := Random(16)
	dst, err := gaes.EncryptCBC([]byte(data), []byte(key), iv)
	if err != nil {
		g.Log().Error("encrypt error", err)
		return []byte{}
	}
	return append(iv, dst...)
}

func Decrypt(key string, data []byte) string {
	iv := data[:16]
	src2 := data[16:]
	dst, err := gaes.DecryptCBC(src2, []byte(key), iv)
	if err != nil {
		g.Log().Error("decrypt error", err)
		return ""
	}
	return string(dst)
}

// GetHash32 encrypts any type of variable using MD5 algorithms.
func GetHash32(content string) string {
	return gmd5.MustEncrypt(content)
}

// Encode  aes CBC 加密.iv可为空(随机生成16位,并在返回的结果里合并到前面)
func Encode(key, iv, data string) string {
	if key == "" {
		return "异常, key不能为空"
	}
	var ivb []byte
	if iv == "" {
		ivb = Random(16)
	} else {
		ivb = []byte(iv)
	}
	dst, err := gaes.EncryptCBC([]byte(data), []byte(key), ivb)
	if err != nil {
		g.Log().Error("encrypt error", err)
		return "异常," + err.Error()
	}
	if iv == "" {
		dst = append(ivb, dst...)
	}

	return base64.StdEncoding.EncodeToString(dst)
}

// Decode aes CBC 解密.iv可为空(从data里解出来,前16位)
func Decode(key, iv, data string) string {
	if key == "" {
		return "异常, key不能为空"
	}
	src2, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		g.Log().Error("DecodeString error", err)
		return "异常," + err.Error()
	}
	var ivb []byte
	if iv == "" {
		ivb = src2[:16]
		src2 = src2[16:]
	} else {
		ivb = []byte(iv)
	}
	dst, err := gaes.DecryptCBC(src2, []byte(key), ivb)
	if err != nil {
		g.Log().Error("decrypt error", err)
		return "异常," + err.Error()
	}
	return string(dst)
}
