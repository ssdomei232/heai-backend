package project

import (
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/model"
)

// 获取项目列表
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

// 通过id获取项目中所有生成任务
func getProjectImageGenerateTasks(projectID int, uid int) (tasks []*model.GenerateTask, err error) {
	// 1.获取数据库连接
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 2.执行查询
	rows, err := db.Query("SELECT id, create_at, finish_at, model, prompt, reference_image_filepaths, category, result_filepath, status, failure_reason, error, sora2_pid FROM generate_task WHERE project_id = ? AND uid = ?", projectID, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 3.扫描Row
	for rows.Next() {
		var task model.GenerateTask
		err := rows.Scan(&task.ID, &task.CreateAt, &task.FinishAt, &task.Model, &task.Prompt, &task.ReferenceImageFilepaths, &task.Category, &task.ResultFilepath, &task.Status, &task.FailureReason, &task.Error, &task.Sora2PID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// 创建项目
func createProject(uid int, createAt int64, title string) error {
	// 1.获取数据库连接
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// 2.执行插入
	_, err = db.Exec("INSERT INTO project (uid, create_at, title) VALUES (?, ?, ?)", uid, createAt, title)
	if err != nil {
		return err
	}

	return nil
}

// 删除项目
func deleteProject(projectID int, uid int) error {
	// 1.获取数据库连接
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// 2.执行删除
	_, err = db.Exec("DELETE FROM project WHERE id = ? AND uid = ?", projectID, uid)
	if err != nil {
		return err
	}

	return nil
}
