package zfile

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
