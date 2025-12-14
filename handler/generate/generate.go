package generate

import (
	"fmt"
	"log"
	"strings"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/configs"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/pkg/grsai"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	NanoBananaProGeneratePointCost  = 10
	NanoBananaFastGeneratePointCost = 3
	NanoBananaGeneratePointCost     = 8
)

func HandleGenerateNanoBanana(c *gin.Context) {
	var err error

	var req NanoBananaGenerateRequest
	err = c.BindJSON(&req)
	// fuck the frontend
	if err = checkNanoBananaGenerateRequest(&req); err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error()})
		return
	}

	// 计费部分
	session := sessions.Default(c)
	username := session.Get("username")
	userInfo, err := user.GetUserInfo(username.(string))
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}
	if userInfo.Point < 0 {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "余额不足",
		})
		return
	}
	switch req.Model {
	case "nano-banana-pro":
		err = CostPoint(userInfo.UID, NanoBananaProGeneratePointCost)
	case "nano-banana-fast":
		err = CostPoint(userInfo.UID, NanoBananaFastGeneratePointCost)
	default:
		err = CostPoint(userInfo.UID, NanoBananaGeneratePointCost)
	}
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "扣费失败"})
		return
	}

	config, err := configs.GetConfig()
	if err != nil {
		log.Printf("加载配置文件失败: %v", err)
		c.JSON(500, gin.H{"error": "加载配置文件失败"})
		return
	}

	fileBase64, err := getInferenceFileBase64(req.FilePaths)
	if err != nil {
		log.Printf("参考图片base64编码列表失败:%v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "参考图片base64编码失败",
		})
		return
	}

	token, err := generateToken()
	if err != nil {
		// if this error happen,maybe the earth has boomd
		c.JSON(500, gin.H{
			"code": 500,
			"data": "世界爆炸了，请稍后再试",
		})
	}
	webhookURL := fmt.Sprintf("https://%s/webhook/nano-banana?t=%s", config.BaseURL, token)

	var resp *grsai.WebhookResult
	grsaiClient := grsai.NewClient(config.GsraiToken)
	resp, err = grsaiClient.NanoBananaGenerateImage(req.Model, req.Prompt, req.AspectRatio, req.ImageSize, fileBase64, webhookURL)
	if err != nil {
		log.Printf("生成图片失败:%v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "生成图片失败",
		})
		return
	}

	if resp.Code != 0 {
		log.Printf("生成图片失败:%v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "生成图片失败",
		})
		return
	}

	err = createNanoBananaGenerateTaskInDB(req.Model, req.Prompt, req.FilePaths, token, resp.Data.ID)
	if err != nil {
		log.Printf("在数据库中创建NanoBanano生成任务失败:%v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "创建NanoBanano生成任务失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": "任务创建成功",
	})
}

// fuck the frontend
func checkNanoBananaGenerateRequest(req *NanoBananaGenerateRequest) error {
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

// 获取参考图片base64编码列表
func getInferenceFileBase64(filepaths []string) (fileBase64 []string, err error) {
	for _, filepath := range filepaths {
		fileB64, err := imageToBase64(filepath)
		if err != nil {
			return nil, err
		}
		fileBase64 = append(fileBase64, fileB64)
	}
	return fileBase64, nil
}

func createNanoBananaGenerateTaskInDB(model string, prompt string, filepaths []string, webhookToken string, dataID string) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var filepathsStr strings.Builder
	for _, filepath := range filepaths {
		filepathsStr.WriteString(filepath + ",")
	}

	_, err = db.Exec("INSERT INTO nanobanana_generate_tasks (model, prompt, reference_image_filepaths, webhook_token, status, data_id) VALUES (?, ?, ?, ?, ?, ?)",
		model, prompt, filepathsStr.String(), webhookToken, "running", dataID)
	if err != nil {
		return err
	}

	return nil
}
