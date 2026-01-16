package file

import "github.com/gin-gonic/gin"

// 获取指定路径的媒体文件
func HandleGetFile(c *gin.Context) {
	filepath := c.Query("f")

	// 权限检查
	if hasPermission(c) {
		c.File(filepath)
		return
	} else {
		c.JSON(403, gin.H{"code": 403, "data": "无权限访问"})
		return
	}
}
