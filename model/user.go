package model

import (
	"fmt"

	"git.mmeiblog.cn/HEntropyAI/HEntropyAI/handler/db"
)

type User struct {
	UID       int     `json:"uid"`                // 用户ID
	Username  string  `json:"username"`           // 用户名
	Password  string  `json:"password,omitempty"` // 密码
	Point     int     `json:"point"`              // 用户积分
	AvatarURL *string `json:"avatar_url"`         // 头像URL
	Email     *string `json:"email"`              // 电子邮箱
	CreateAt  int64   `json:"create_at"`          // 创建时间戳
}

type ImageGenerateTask struct {
	ID                      int     `json:"id"`
	UID                     int     `json:"uid"`
	DataID                  *string `json:"data_id"`
	Model                   string  `json:"model"`
	Prompt                  string  `json:"prompt"`
	ReferenceImageFilepaths *string `json:"reference_image_filepaths"`
	ResultURL               *string `json:"result_url"`
	ResultFilepath          *string `json:"result_filepath"`
	Status                  string  `json:"status"`
	FailureReason           *string `json:"failure_reason"`
	Error                   *string `json:"error"`
}

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
