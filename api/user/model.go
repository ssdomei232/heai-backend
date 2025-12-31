package user

type User struct {
	UID       int     `json:"uid"`                // 用户ID
	Username  string  `json:"username"`           // 用户名
	Password  string  `json:"password,omitempty"` // 密码
	Point     int     `json:"point"`              // 用户积分
	AvatarURL *string `json:"avatar_url"`         // 头像URL
	Email     *string `json:"email"`              // 电子邮箱
	CreateAt  int64   `json:"create_at"`          // 创建时间戳
}

type NanoBananaGenerateTask struct {
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
