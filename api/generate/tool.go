package generate

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
)

func imageToBase64(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func CostPoint(userID int, point int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE user SET point = point - ? WHERE uid = ?", point, userID)
	return err
}
