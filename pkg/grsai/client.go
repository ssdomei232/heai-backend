package grsai

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(token string) *Client {
	restyClient := resty.New().
		SetRetryCount(3).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() >= 500 // 仅在服务端返回5xx时重试
			})
	restyClient.SetBaseURL("https://grsai.dakka.com.cn")
	restyClient.SetHeader("Authorization", "Bearer "+token)
	restyClient.SetHeader("Accept", "application/json")

	return &Client{
		client: restyClient,
	}
}

// DoRequest 发起一次http请求
func (c *Client) DoRequest(method, endpoint string, reqData any, respData any) error {
	req := c.client.R()

	if reqData != nil {
		req.SetBody(reqData)
	}

	if respData != nil {
		req.SetResult(respData)
	}

	_, err := req.Execute(method, endpoint)

	// 错误处理
	return err
}
