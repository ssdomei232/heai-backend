package sora2

import (
	"fmt"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
)

func checkSora2GenerateRequest(req *model.VideoGenerateRequest) error {
	var validModels = map[string]struct{}{
		"sora-2": {},
	}
	var validAspectRatios = map[string]struct{}{
		"16:9": {}, "9:16": {},
	}
	var validSizes = map[string]struct{}{
		"large": {}, "small": {},
	}
	var validDurations = map[int]struct{}{
		10: {}, 15: {},
	}

	if _, ok := validDurations[req.Duration]; !ok {
		return fmt.Errorf("不支持的时长")
	}
	if _, ok := validModels[req.Model]; !ok {
		return fmt.Errorf("不支持的模型")
	}
	if _, ok := validAspectRatios[req.AspectRatio]; !ok {
		return fmt.Errorf("不支持的比例")
	}
	if _, ok := validSizes[req.Size]; !ok {
		return fmt.Errorf("不支持的大小")
	}
	return nil
}

// 在数据库中创建视频生成任务记录
func createGenerateTaskInDB(uid int, model string, prompt string, filepath string, webhookToken string, dataID string, ProjectID int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO sora2_generate_task (uid, model, create_at, prompt, reference_image_filepath, webhook_token, status, data_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		uid, model, time.Now().Unix(), prompt, filepath, webhookToken, "running", dataID)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO generate_task (uid, create_at, model, prompt, webhook_token, reference_image_filepaths, category, status, project_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		uid, time.Now().Unix(), model, prompt, webhookToken, filepath, "video", "running", ProjectID)
	if err != nil {
		return err
	}

	return nil
}
