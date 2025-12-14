package user

type User struct {
	UID       int     `json:"uid"`
	Username  string  `json:"username"`
	Password  string  `json:"password,omitempty"`
	Point     int     `json:"point"`
	AvatarURL *string `json:"avatar_url"`
	Email     *string `json:"email"`
	CreateAt  int64   `json:"create_at"`
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
