package upload

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"github.com/gin-gonic/gin"
)

// 上传文件
func HandleUploadImage(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "data": "上传文件错误"})
		return
	}

	// 验证文件类型是否为图片
	if !isImageFile(file.Filename) {
		c.JSON(400, gin.H{"code": 400, "data": "只允许上传图片文件"})
		return
	}

	// 1. 生成日期目录结构
	now := time.Now()
	dateDir := "data/" + time.Now().Format("2006/01/02") + "/"

	// 2.生成基于文件名和时间戳的哈希值
	hashFileName := generateHashFileName(file.Filename, now.Unix())
	// 3.获取文件扩展名
	ext := filepath.Ext(file.Filename)
	// 4.组合完整路径
	filePath := filepath.Join(dateDir, hashFileName+ext)

	// 5.确保目录存在
	if err := os.MkdirAll(dateDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"code": 500, "data": "创建目录失败"})
		return
	}

	// 6.保存文件
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": "保存文件失败"})
		return
	}

	// 7.写入数据库
	err = writeToDatabase(userInfo.UID, filePath)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": "服务器内部错误"})
		return
	}

	c.JSON(200, gin.H{
		"code":     200,
		"data":     "上传成功",
		"filePath": filePath,
	})
}

func HandleGetUploadedFiles(c *gin.Context) {
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	files, err := getUploadFiles(userInfo.UID)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取上传文件列表失败"})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": files,
	})
}
