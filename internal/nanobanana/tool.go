package nanobanana

import (
	"fmt"
	"strings"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
)

// 校验前端参数是否合法
func checkNanoBananaGenerateRequest(req *model.ImageGenerateRequest) error {
	var validModels = map[string]struct{}{
		"nano-banana-fast":   {},
		"nano-banana":        {},
		"nano-banana-pro":    {},
		"nano-banana-pro-vt": {},
	}
	var validAspectRatios = map[string]struct{}{
		"1:1": {}, "16:9": {}, "9:16": {}, "4:3": {}, "3:4": {},
		"3:2": {}, "2:3": {}, "5:4": {}, "4:5": {}, "21:9": {}, "auto": {},
	}
	var validImageSizes = map[string]struct{}{
		"1K": {}, "2K": {}, "4K": {},
	}
	if _, ok := validModels[req.Model]; !ok {
		return fmt.Errorf("不支持的模型")
	}
	if _, ok := validAspectRatios[req.AspectRatio]; !ok {
		return fmt.Errorf("不支持的比例")
	}
	if _, ok := validImageSizes[req.ImageSize]; !ok {
		return fmt.Errorf("不支持的图片大小")
	}
	return nil
}

// 在数据库中创建图像生成任务记录
func createGenerateTaskInDB(uid int, model string, prompt string, filepaths []string, webhookToken string, dataID string, groupID int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var filepathsStr strings.Builder
	for _, filepath := range filepaths {
		filepathsStr.WriteString(filepath + ",")
	}

	_, err = db.Exec("INSERT INTO nanobanana_generate_task (uid, model, create_at, prompt, reference_image_filepaths, webhook_token, status, data_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		uid, model, time.Now().Unix(), prompt, filepathsStr.String(), webhookToken, "running", dataID)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO generate_task (uid, create_at, model, prompt, webhook_token, reference_image_filepaths, category, status, project_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		uid, time.Now().Unix(), model, prompt, webhookToken, filepathsStr.String(), "image", "running", groupID)
	if err != nil {
		return err
	}

	return nil
}
