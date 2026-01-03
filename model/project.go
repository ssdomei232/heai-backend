package model

type Project struct {
	ID       int    `json:"id"`
	UID      int    `json:"uid"`
	CreateAT int    `json:"create_at"`
	Title    string `json:"title"`
}

type GenerateTask struct {
	ID                      int     `json:"id"`
	CreateAt                int     `json:"create_at"`
	FinishAt                int     `json:"finish_at"`
	Model                   string  `json:"model"`
	Prompt                  string  `json:"prompt"`
	ReferenceImageFilepaths *string `json:"reference_image_filepaths"`
	Category                string  `json:"category"` // video or image
	ResultFilepath          *string `json:"result_filepath"`
	Status                  string  `json:"status"`
	FailureReason           *string `json:"failure_reason"`
	Error                   *string `json:"error"`
}
