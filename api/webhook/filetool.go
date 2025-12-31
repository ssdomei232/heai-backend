package webhook

import "github.com/imroc/req/v3"

func downloadFile(url string, filepath string) error {
	_, err := req.R().SetOutputFile(filepath).Get(url)
	return err
}
