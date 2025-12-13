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
