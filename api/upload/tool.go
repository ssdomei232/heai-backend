package upload

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
)

// 验证是否为图片文件
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}

	return slices.Contains(allowedExtensions, ext)
}

// 生成基于文件名和时间戳的哈希文件名
func generateHashFileName(originalName string, timestamp int64) string {
	data := fmt.Sprintf("%s%d", originalName, timestamp)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	return hash
}

// 写入数据库
func writeToDatabase(uid int, filepath string) error {
	// 创建数据库连接
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// 插入数据
	_, err = db.Exec("INSERT INTO upload_image (uid, filepath) VALUES (?, ?)", uid, filepath)
	if err != nil {
		return err
	}

	return nil
}

// 获取上传的文件
func getUploadFiles(uid int) (uploadFiles []string, err error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT filepath FROM upload_image WHERE uid = ?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var filepath string
		err = rows.Scan(&filepath)
		if err != nil {
			return nil, err
		}
		uploadFiles = append(uploadFiles, filepath)
	}

	return uploadFiles, nil
}
