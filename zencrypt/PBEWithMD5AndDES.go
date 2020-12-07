package zencrypt

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"strings"
)

type PBEWithMD5AndDES struct {
	iterations int
	password   string
}

var PBE = &PBEWithMD5AndDES{1000, ""}

func (pbe *PBEWithMD5AndDES) getDerivedKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count-1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

func (pbe *PBEWithMD5AndDES) encrypt(password string, obtenationIterations int, plainText string, salt []byte) (string, error) {
	padNum := byte(8 - len(plainText)%8)
	for i := byte(0); i < padNum; i++ {
		plainText += string(padNum)
	}

	dk, iv := pbe.getDerivedKey(password, salt, obtenationIterations)

	block, err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plainText))
	encrypter.CryptBlocks(encrypted, []byte(plainText))

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (pbe *PBEWithMD5AndDES) decrypt(password string, obtenationIterations int, cipherText string, salt []byte) (string, error) {
	msgBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	dk, iv := pbe.getDerivedKey(password, salt, obtenationIterations)
	block, err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(msgBytes))
	decrypter.CryptBlocks(decrypted, msgBytes)

	decryptedString := strings.TrimRight(string(decrypted), "\x01\x02\x03\x04\x05\x06\x07\x08")

	return decryptedString, nil
}

//func test() {
//	salt := []byte{0xFF, 0x2B, 0x38, 0x30, 0xF8, 0x61, 0xEF, 0x99}
//	password := "my_secret_password"
//	iterations := 222
//	originalText := "mythings"
//
//	res, err := encrypt(password, iterations, originalText, salt)
//	fmt.Println("encripted", res, err)
//	res, err = decrypt(password, iterations, res, salt)
//	fmt.Println("decripted", res, err)
//}

func (pbe *PBEWithMD5AndDES) Encrypt(salt, content string) (string, error) {
	return pbe.encrypt("", 1000, content, ([]byte)(salt))
}
func (pbe *PBEWithMD5AndDES) Decrypt(salt, content string) (string, error) {
	return pbe.decrypt("", 1000, content, ([]byte)(salt))
}
