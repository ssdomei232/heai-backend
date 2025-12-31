package user

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleRegistry(c *gin.Context) {
	var user User
	var err error

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "data": "输入错误"})
		return
	}

	if err = user.IsValid(); err != nil {
		c.JSON(400, gin.H{"code": 400, "data": err.Error()})
		return
	}

	if isExist := user.IsExist(); isExist {
		c.JSON(400, gin.H{"code": 400, "data": "用户名已存在"})
		return
	}

	if err = createUser(&user); err != nil {
		c.JSON(500, gin.H{"code": 500, "data": "注册失败"})
		return
	}

	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Save()

	c.JSON(200, gin.H{"code": 200, "data": "注册成功"})
}

func HandleLogin(c *gin.Context) {
	var user User
	var err error

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "data": "输入错误"})
		return
	}

	if err = user.IsValid(); err != nil {
		c.JSON(400, gin.H{"code": 400, "data": err.Error()})
	}

	if err = verifyUser(&user); err != nil {
		c.JSON(400, gin.H{"code": 400, "data": "用户名或密码错误"})
		return
	}

	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Save()

	c.JSON(200, gin.H{"code": 200, "data": "登录成功"})
}

func HandleGetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	userInfo, err := GetUserInfo(username.(string))
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": userInfo})
}

func HandleGetNanoBananaGenerateTask(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	userInfo, err := GetUserInfo(username.(string))
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取用户信息失败"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"code": 400, "data": "参数错误"})
		return
	}
	perpage, err := strconv.Atoi(c.DefaultQuery("perpage", "10"))
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"code": 400, "data": "参数错误"})
		return
	}
	tasks, allRecords, err := GetNanoBananaGenerateTask(userInfo.UID, page, perpage)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"code": 500, "data": "获取任务失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"all_records": allRecords, "tasks": tasks}})
}
