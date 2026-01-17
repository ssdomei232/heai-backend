package model

type ImageGenerateRequest struct {
	Prompt      string   `json:"prompt"`
	Model       string   `json:"model"`
	AspectRatio string   `json:"aspectRatio"` // 比例
	ImageSize   string   `json:"imageSize,omitempty"`
	FilePaths   []string `json:"filepaths"`  // 参考图片路径列表
	ProjectID   int      `json:"project_id"` // 项目ID
}

type VideoGenerateRequest struct {
	Prompt        string `json:"prompt"`
	Model         string `json:"model"`
	AspectRatio   string `json:"aspectRatio"` // 比例
	Duration      int    `json:"duration"`    // 视频时长
	Filepath      string `json:"filepath"`    // 参考图片路径
	RemixTargetID string `json:"remixTargetID,omitempty"`
	Size          string `json:"size,omitempty"` // 视频清晰度 small/large
	ProjectID     int    `json:"project_id"`     // 项目ID
}
