package grsai

// NanoBanana 图片生成请求参数
type NanoBananaRequest struct {
	Model        string   `json:"model"`       // 模型名称
	Prompt       string   `json:"prompt"`      // 文本提示词
	AspectRatio  string   `json:"aspectRatio"` // 比例
	ImageSize    string   `json:"imageSize"`
	Urls         []string `json:"urls"` // Base64
	WebHook      string   `json:"webHook"`
	ShutProgress bool     `json:"shutProgress"`
}

// Webhook 结果
type WebhookResult struct {
	Code int    `json:"code"` // 0 表示成功
	Msg  string `json:"msg"`  // success
	Data struct {
		ID string `json:"id"` // id
	} `json:"data"`
}

// NanoBanana webHook 响应参数
type NanoBananaWebhook struct {
	ID      string `json:"id"` // id
	Results []struct {
		URL     string `json:"url"`     // 生成图片的URL
		Content string `json:"content"` // 模型回复的文本内容
	} `json:"results"`
	Progress      int    `json:"progress"` // 进度
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason"`
	Error         string `json:"error"`
}

// Sora2 视频生成请求参数
type Sora2Request struct {
	Model         string `json:"model"`
	Prompt        string `json:"prompt"`
	URL           string `json:"url"`
	AspectRatio   string `json:"aspectRatio"`
	Duration      int    `json:"duration"`
	RemixTargetID string `json:"remixTargetId"`
	Size          string `json:"size"`
	WebHook       string `json:"webHook"`
	ShutProgress  bool   `json:"shutProgress"`
}

// Sora2 webHook 响应参数
type Sora2Webhook struct {
	ID      string `json:"id"`
	Results []struct {
		URL             string `json:"url"`             // 视频URL
		RemoveWatermark bool   `json:"removeWatermark"` // 是否去水印
		Pid             string `json:"pid"`             // pid: 视频续作, remix目标id
	} `json:"results"`
	Progress      int    `json:"progress"`       // 进度
	Status        string `json:"status"`         // 任务状态 "running": 进行中 "succeeded": 成功 "failed": 失败
	FailureReason string `json:"failure_reason"` // 失败原因 "output_moderation": 输出违规 "input_moderation": 输入违规 "error": 其他错误
	Error         string `json:"error"`          // 失败详细信息
}
