package file

import (
	"log"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"github.com/gin-gonic/gin"
)

// 检查用户是否有权限访问媒体文件
func hasPermission(c *gin.Context) bool {
	// 1. 获取用户信息
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		return false
	}

	// 2. 连接数据库
	db, err := db.GetDB()
	if err != nil {
		return false
	}
	defer db.Close()

	// 3. 在数据库中查询用户是否有该filepath的访问权限
	filepath := c.Query("f")
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM generate_task WHERE uid = ? AND result_filepath = ?", userInfo.UID, filepath).Scan(&count)
	log.Println(userInfo.UID, filepath, count)
	if err != nil {
		return false
	}

	// 4. 如果count大于0，说明有权限
	return count > 0
}
