package generate

type NanoBananaGenerateRequest struct {
	Prompt      string   `json:"prompt"`
	Model       string   `json:"model"`
	AspectRatio string   `json:"aspectRatio"` // 比例
	ImageSize   string   `json:"imageSize"`
	FilePaths   []string `json:"filepaths"` // 参考图片路径列表
}
