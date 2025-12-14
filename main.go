package main

import (
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/generate"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/webhook"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/internal/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("Ahu0eM3xlOvxEJiwa"))
	r.Use(sessions.Sessions("cat-session", store))

	{
		v1 := r.Group("/v1")
		// 不需要认证的路由
		v1.POST("/user/registry", user.HandleRegistry)
		v1.POST("/user/login", user.HandleLogin)
		v1.POST("/webhook/nano-banana", webhook.HandleNanoBananaWebhook)

		// 需要认证的路由组
		authorized := v1.Group("/")
		authorized.Use(AuthMiddleware())
		{
			authorized.POST("/generate/nano-banana", generate.HandleGenerateNanoBanana)
			authorized.GET("/user/info", user.HandleGetUserInfo)
			authorized.GET("/user/nano-banana-task", user.HandleGetNanoBananaGenerateTask)
		}
	}

	r.Run(":8077")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.JSON(401, gin.H{"code": 401, "data": "未登录"})
			c.Abort()
			return
		}
		c.Next()
	}
}
