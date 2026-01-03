package webhook

import (
	"log"
	"os"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/pkg/grsai"
	"github.com/gin-gonic/gin"
)

func HandleNanoBananaWebhook(c *gin.Context) {
	var token string
	var resp grsai.NanoBananaWebhook
	var err error

	token = c.DefaultQuery("t", "") // Get token
	err = c.BindJSON(&resp)
	if err != nil {
		log.Printf("webhook回调时绑定json出错: %v", err)
		c.Status(200)
		return
	}

	// 如果生成成功，下载图片并更新任务状态
	if resp.Status == "succeeded" && len(resp.Results) > 0 {
		err = os.MkdirAll("data/"+time.Now().Format("2006/01/02"), 0755)
		if err != nil {
			log.Printf("创建目录失败: %v", err)
			return
		}

		filepath := "data/" + time.Now().Format("2006/01/02") + "/" + resp.ID + ".png"
		err = downloadFile(resp.Results[0].URL, filepath)
		if err != nil {
			log.Printf("下载NanoBanana生成的图片失败: %v", err)
			return
		}

		err = updateNanoBananaTaskStatus(token, resp.Status, resp.Results[0].URL, filepath, "", "")
		if err != nil {
			log.Printf("更新NanoBanana任务状态失败: %v", err)
			return
		}
		c.Status(200)
		return
	} else {
		err = updateNanoBananaTaskStatus(token, resp.Status, "", "", resp.FailureReason, resp.Error)
		if err != nil {
			log.Printf("更新NanoBanana任务状态失败: %v", err)
			return
		}
		c.Status(200)
		return
	}
}

func updateNanoBananaTaskStatus(webhookToken string, status string, resultURL string, resultFilepath string, failureReason string, error string) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE nanobanana_generate_task SET status = ?, result_url = ?, result_filepath = ?, failure_reason = ?, error = ?, finish_at = ? WHERE webhook_token = ?",
		status, resultURL, resultFilepath, failureReason, error, time.Now().Unix(), webhookToken)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE generate_task SET status = ?, result_filepath = ?, failure_reason = ?, error = ?, finish_at = ? WHERE webhook_token = ?",
		status, resultFilepath, failureReason, error, time.Now().Unix(), webhookToken)
	if err != nil {
		return err
	}

	return nil
}
