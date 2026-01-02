package generate

import (
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
)

func CostPoint(userID int, point int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE user SET point = point - ? WHERE uid = ?", point, userID)
	return err
}
