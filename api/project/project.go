package project

import (
	"log"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/api/user"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleGetProjects(c *gin.Context) {
	var err error
	session := sessions.Default(c)
	username := session.Get("username")
	userInfo, err := user.GetUserInfo(username.(string))
	if err != nil {
		log.Println(err)
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

func getProjects(uid int) (projects []*model.Project, err error) {
	// 1.获取数据库连接
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 2.执行查询
	rows, err := db.Query("SELECT id, title, create_at FROM project WHERE uid = ?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project model.Project
		err = rows.Scan(&project.ID, &project.Title, &project.CreateAT)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}
