package model

type ImageGenerateRequest struct {
	Prompt      string   `json:"prompt"`
	Model       string   `json:"model"`
	AspectRatio string   `json:"aspectRatio"` // 比例
	ImageSize   string   `json:"imageSize,omitempty"`
	FilePaths   []string `json:"filepaths"`  // 参考图片路径列表
	ProjectID   int      `json:"project_id"` // 项目ID
}
