package zfile

import (
	"encoding/json"
	"github.com/zhang201702/zhang/zlog"
	"io/ioutil"
	"os"
	"strings"
)

// OpenTest 打开一个文本文件
func OpenText(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func OpenJson(filePath string, obj interface{}) error {
	text, err := OpenText(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(text), obj)
}

func AppendText(filePath, text string) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		zlog.LogError(err, "zfile.AppendText 异常", filePath)
		return
	}
	defer f.Close()
	f.WriteString(strings.TrimSpace("") + text)
}
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
