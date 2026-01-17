package tool

import (
	"encoding/base64"
	"os"
)

// 获取参考图片base64编码列表
func GetFileBase64(filepaths []string) (fileBase64 []string, err error) {
	for _, filepath := range filepaths {
		fileB64, err := ImageToBase64(filepath)
		if err != nil {
			return nil, err
		}
		fileBase64 = append(fileBase64, fileB64)
	}
	return fileBase64, nil
}

func ImageToBase64(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}
