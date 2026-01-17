package generate

import (
	"html/template"
	"log"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/nanobanana"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
	"github.com/gin-gonic/gin"
)

const (
	NanoBananaProGeneratePointCost  = 50
	NanoBananaFastGeneratePointCost = 10
	NanoBananaGeneratePointCost     = 20
)

func HandleGenerateImage(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	// 1.解析请求参数
	var req model.ImageGenerateRequest
	if err = c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "JSON 解析失败"})
		log.Println(err)
		return
	}
	req.Prompt = template.HTMLEscapeString(req.Prompt) // 防止 XSS 攻击

	// 2.计费部分
	var pointcost int
	switch req.Model {
	case "nano-banana-pro":
		pointcost = NanoBananaProGeneratePointCost
	case "nano-banana-fast":
		pointcost = NanoBananaFastGeneratePointCost
	default:
		pointcost = NanoBananaGeneratePointCost
	}
	if userInfo.Point < pointcost {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "余额不足",
		})
		return
	}
	err = CostPoint(userInfo.UID, pointcost)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "扣费失败"})
		return
	}

	// 3.调用生成函数
	errorCode, err := nanobanana.Generate(&req, userInfo)
	if err != nil {
		c.JSON(500, gin.H{
			"code": errorCode,
			"data": err.Error(),
		})
		return
	}

	// 4.返回成功响应
	c.JSON(200, gin.H{
		"code": 200,
		"data": "任务创建成功",
	})
}
