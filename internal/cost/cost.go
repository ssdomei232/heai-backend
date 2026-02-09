package cost

import (
	"fmt"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
)

func CostPointsByModelName(uid int, userPoint int, modelName string) error {
	var pointcost int
	switch modelName {
	case "sora-2":
		pointcost = Sora2GeneratePointCost
	case "nano-banana-pro":
		pointcost = NanoBananaProGeneratePointCost
	case "nano-banana-fast":
		pointcost = NanoBananaFastGeneratePointCost
	case "nano-banana":
		pointcost = NanoBananaGeneratePointCost
	default:
		pointcost = 0
	}

	if pointcost == 0 {
		return fmt.Errorf("错误的模型名称")
	}

	if userPoint < pointcost {
		return fmt.Errorf("余额不足")
	}

	return costPoint(uid, pointcost)
}

func costPoint(uid int, pointcost int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET point = point - ? WHERE uid = ?", pointcost, uid)
	return err
}
