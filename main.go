package main

import (
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/generate"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/webhook"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	{
		v1 := r.Group("/v1")
		v1.POST("/webhook/nano-banana", webhook.HandleNanoBananaWebhook)

		v1.POST("/generate/nano-banana", generate.HandleGenerateNanoBanana)
	}
}
