package grsai

func (c *Client) NanoBananaGenerateImage(model string, prompt string, aspectRatio string, imageSize string, urls []string, webHook string, shutProgress bool) (*WebhookResult, error) {
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
