package grsai

// 生成 NanoBanana 图片
//
// model： 模型名称
//
// prompt： 文本提示词
//
// aspectRatio： 比例
//
// imageSize： 图片尺寸
//
// urls： base64编码的图片列表
func (c *Client) NanoBananaGenerateImage(model string, prompt string, aspectRatio string, imageSize string, urls []string, webHook string) (*WebhookResult, error) {
	reqData := &NanoBananaRequest{
		Model:        model,
		Prompt:       prompt,
		AspectRatio:  aspectRatio,
		ImageSize:    imageSize,
		Urls:         urls,
		WebHook:      webHook,
		ShutProgress: true,
	}

	var respData WebhookResult
	err := c.DoRequest("POST", "/v1/draw/nano-banana", reqData, &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}

// 生成 Sora2 视频
//
// model： 模型名称
//
// prompt： 提示词
//
// url： 参考图片URL
//
// aspectRatio： 宽高比 9:16、16:9
//
// duration： 视频时长 10/15
//
// remixTargetID： remix目标id 可参考返回结果的pid值: s_xxxxxxxxx
//
// size： 视频清晰度 small/large
//
// webHook： webHook地址
func (c *Client) Sora2GenerateVideo(model string, prompt string, aspectRatio string, url string, duration int, remixTargetID string, size string, webHook string) (*WebhookResult, error) {
	reqData := &Sora2Request{
		Model:         model,
		Prompt:        prompt,
		URL:           url,
		AspectRatio:   aspectRatio,
		Duration:      duration,
		RemixTargetID: remixTargetID,
		Size:          size,
		WebHook:       webHook,
		ShutProgress:  true,
	}

	var respData WebhookResult
	err := c.DoRequest("POST", "/v1/video/sora-video", reqData, &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}
