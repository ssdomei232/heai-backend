package csrf

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

// 适配 Gorilla CSRF 中间件到 Gin
func GinCSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 gorilla CSRF 中间件包装当前的处理函数
		csrfMiddleware := csrf.Protect([]byte("Ahu0eM3xlOvxEJiwa"),
			csrf.Secure(true), // 在开发环境中设置为false，生产环境中应为true
			csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 将 http.ResponseWriter 和 http.Request 转换回 gin.Context
				c.JSON(403, gin.H{"code": 403, "data": "CSRF token mismatch"})
				c.Abort()
			})),
		)

		// 应用 CSRF 中间件
		handler := csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 保存修改后的请求到 gin.Context
			c.Request = r
			c.Next()
		}))

		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func HandleGetCSRFToken(c *gin.Context) {
	token := csrf.Token(c.Request)
	c.JSON(200, gin.H{"code": 200, "data": token})
}
