package webhook

import (
	"log"
	"os"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/pkg/grsai"
	"github.com/gin-gonic/gin"
)

func HandleSora2Webhook(c *gin.Context) {
	var token string
	var resp grsai.Sora2Webhook
	var err error

	token = c.DefaultQuery("t", "") // Get token
	err = c.BindJSON(&resp)
	if err != nil {
		log.Printf("webhook回调时绑定json出错: %v", err)
		c.Status(200)
		return
	}

	// 如果生成成功，下载视频并更新任务状态
	if resp.Status == "succeeded" && len(resp.Results) > 0 {
		err = os.MkdirAll("data/"+time.Now().Format("2006/01/02"), 0755)
		if err != nil {
			log.Printf("创建目录失败: %v", err)
			return
		}

		filepath := "data/" + time.Now().Format("2006/01/02") + "/" + resp.ID + ".png"
		err = downloadFile(resp.Results[0].URL, filepath)
		if err != nil {
			log.Printf("下载Sora2生成的视频失败: %v", err)
			return
		}

		err = updateSora2TaskStatus(token, resp.Status, resp.Results[0].URL, filepath, "", "", resp.Results[0].Pid)
		if err != nil {
			log.Printf("更新Sora2任务状态失败: %v", err)
			return
		}
		c.Status(200)
		return
	} else {
		err = updateSora2TaskStatus(token, resp.Status, "", "", resp.FailureReason, resp.Error, resp.Results[0].Pid)
		if err != nil {
			log.Printf("更新Sora2任务状态失败: %v", err)
			return
		}
		c.Status(200)
		return
	}
}

func updateSora2TaskStatus(webhookToken string, status string, resultURL string, resultFilepath string, failureReason string, error string, sora2Pid string) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE sora2_generate_task SET status = ?, result_url = ?, result_filepath = ?, failure_reason = ?, error = ?, finish_at = ?, sora2_pid = ? WHERE webhook_token = ?",
		status, resultURL, resultFilepath, failureReason, error, time.Now().Unix(), sora2Pid, webhookToken)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE generate_task SET status = ?, result_filepath = ?, failure_reason = ?, error = ?, finish_at = ?, sora2_pid = ? WHERE webhook_token = ?",
		status, resultFilepath, failureReason, error, time.Now().Unix(), sora2Pid, webhookToken)
	if err != nil {
		return err
	}

	return nil
}
