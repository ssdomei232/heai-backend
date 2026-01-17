package sora2

import (
	"fmt"
	"log"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/configs"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/errorcode"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/tool"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/pkg/grsai"
)

func Generate(req *model.VideoGenerateRequest, userInfo *model.User) (errorCode errorcode.ErrorCode, err error) {
	// 1.验证请求参数
	if err = checkSora2GenerateRequest(req); err != nil {
		return errorcode.CodeInvalidRequest, err
	}

	// 2.加载配置文件
	config, err := configs.GetConfig()
	if err != nil {
		log.Printf("加载配置文件失败: %v", err)
		return errorcode.CodeInternalError, err
	}

	// 3.准备参考图片的base64编码列表
	var fileBase64 string
	if req.Filepath != "" {
		fileBase64, err = tool.ImageToBase64(req.Filepath)
		if err != nil {
			log.Printf("参考图片base64编码列表失败:%v", err)
			return errorcode.CodeInternalError, err
		}
	} else {
		fileBase64 = ""
	}

	// 4.生成webhook token
	token, err := tool.GenerateToken()
	if err != nil {
		return errorcode.CodeInternalError, err
	}
	webhookURL := fmt.Sprintf("https://%s/v1/webhook/sora2?t=%s", config.BaseURL, token)

	// 5.调用GRSAI接口生成图片
	var resp *grsai.WebhookResult
	grsaiClient := grsai.NewClient(config.GsraiToken)
	resp, err = grsaiClient.Sora2GenerateVideo(req.Model, req.Prompt, req.AspectRatio, fileBase64, req.Duration, req.RemixTargetID, req.Size, webhookURL)
	if err != nil {
		log.Printf("生成视频失败:%v", err)
		return errorcode.CodeGenerateImageFailed, fmt.Errorf("生成视频失败")
	}
	if resp.Code != 0 {
		log.Printf("接口异常:%v", err)
		return errorcode.CodeUpstreamAPIError, fmt.Errorf("接口异常,请等待修复")
	}

	// 6.写入数据库
	err = createGenerateTaskInDB(userInfo.UID, req.Model, req.Prompt, req.Filepath, token, resp.Data.ID, req.ProjectID)
	if err != nil {
		log.Printf("在数据库中创建Sora2生成任务失败:%v", err)
		return errorcode.CodeInternalError, fmt.Errorf("创建Sora2生成任务失败")
	}

	return errorcode.CodeSuccess, nil
}
