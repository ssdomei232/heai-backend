package user

import (
	"fmt"
	"time"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) IsValid() error {
	if u.Username == "" || u.Password == "" {
		return fmt.Errorf("用户名或密码不能为空")
	}
	if len(u.Username) > 32 || len(u.Username) < 2 {
		return fmt.Errorf("用户名过长或过短")
	}
	if len(u.Password) > 64 || len(u.Password) < 6 {
		return fmt.Errorf("密码过长或过短")
	}
	return nil
}

func (u *User) IsExist() bool {
	db, err := db.GetDB()
	if err != nil {
		return false
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM user WHERE name = ?", u.Username).Scan(&count)
	return count > 0
}

// 加密密码
func encryptPassword(password string) (string, error) {
	hashedID, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedID), nil
}

// 验证密码
func verifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

func createUser(u *User) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hashedPassword, err := encryptPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO user (name, password, create_time) VALUES (?, ?, ?)", u.Username, hashedPassword, time.Now().Unix())
	return err
}

func verifyUser(u *User) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var hashedPassword string
	err = db.QueryRow("SELECT password FROM user WHERE name = ?", u.Username).Scan(&hashedPassword)
	if err != nil {
		return err
	}

	return verifyPassword(hashedPassword, u.Password)
}

func GetUserInfo(username string) (*User, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT uid, name, point, avatar_url, email, create_time FROM user WHERE name = ?", username).Scan(
		&user.UID, &user.Username, &user.Point, &user.AvatarURL, &user.Email, &user.CreateAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetNanoBananaGenerateTask(uid int, page int, perpage int) (tasks []*NanoBananaGenerateTask, allRecords int, err error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, 0, err
	}
	defer db.Close()

	// 获取总记录数
	err = db.QueryRow("SELECT COUNT(*) FROM nanobanana_generate_task WHERE uid = ?", uid).Scan(&allRecords)
	if err != nil {
		return nil, 0, err
	}

	// 分页查询数据
	offset := (page - 1) * perpage
	rows, err := db.Query("SELECT id, uid, data_id, model, prompt, reference_image_filepaths, result_url, result_filepath, status, failure_reason, error FROM nanobanana_generate_task WHERE uid = ? ORDER BY id DESC LIMIT ? OFFSET ?", uid, perpage, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var task NanoBananaGenerateTask
		err := rows.Scan(&task.ID, &task.UID, &task.DataID, &task.Model, &task.Prompt, &task.ReferenceImageFilepaths, &task.ResultURL, &task.ResultFilepath, &task.Status, &task.FailureReason, &task.Error)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, allRecords, nil
}
