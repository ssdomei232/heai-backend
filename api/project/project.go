package project

import (
	"log"
	"strconv"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
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
	projectID, err := strconv.Atoi(c.Param("id"))
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

// 创建项目
func HandleCreateProject(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	// 1.解析请求参数
	var req model.CreateProjectRequest
	if err = c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "JSON 解析失败"})
		log.Println(err)
		return
	}

	// 2.创建项目
	err = createProject(userInfo.UID, time.Now().Unix(), req.Title)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "data": "创建项目失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "data": "创建项目成功"})
}

// 删除项目
func HandleDeleteProject(c *gin.Context) {
	var err error
	userInfo, err := user.GetUserInfoByGinCtx(c)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"code": 400, "data": "参数错误"})
		return
	}

	err = deleteProject(projectID, userInfo.UID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "data": "删除项目失败"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": "删除项目成功"})
}
