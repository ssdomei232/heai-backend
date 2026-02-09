package generate

import (
	"html/template"
	"log"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/cost"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/nanobanana"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
	"github.com/gin-gonic/gin"
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

	// 2.调用生成函数
	errorCode, err := nanobanana.Generate(&req, userInfo)
	if err != nil {
		c.JSON(500, gin.H{
			"code": errorCode,
			"data": err.Error(),
		})
		return
	}

	// 3.计费部分
	err = cost.CostPointsByModelName(userInfo.UID, userInfo.Point, req.Model)
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"code": 400, "data": err})
		return
	}

	// 4.返回成功响应
	c.JSON(200, gin.H{
		"code": 200,
		"data": "任务创建成功",
	})
}
