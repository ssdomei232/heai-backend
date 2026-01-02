package main

import (
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/csrf"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/generate"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/webhook"
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
		authorized.Use(csrf.GinCSRFMiddleware())
		{
			authorized.POST("/generate/image", generate.HandleGenerateImage)
			authorized.GET("/user/image-task", user.HandleGetImageGenerateTask)
			authorized.GET("/csrf-token", csrf.HandleGetCSRFToken)
			authorized.GET("/user/info", user.HandleGetUserInfo)
		}
	}

	r.Run(":8077")
}

// 认证中间件
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
