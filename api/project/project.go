package project

import (
	"log"
	"strconv"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"github.com/gin-gonic/gin"
)

// 获取所有项目
func HandleGetProjects(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	// 1.获取项目列表
	projects, err := getProjects(userInfo.UID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取项目列表失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"projects": projects}})
}

// 获取指定项目下所有生成任务
func HandleGetProjectDetails(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	// 1.获取项目ID
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"code": 400, "data": "参数错误"})
		return
	}

	// 2.获取项目下所有生成任务
	tasks, err := getProjectImageGenerateTasks(projectID, userInfo.UID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取任务列表失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"tasks": tasks}})
}
