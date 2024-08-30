package webhook

import (
	"fmt"
	"net/http"
	"net/url"
)

func Send(webhookURL string, msg string) error {
	form := url.Values{}
	form.Add("content", msg)

	resp, err := http.PostForm(webhookURL, form)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if http.StatusOK > resp.StatusCode || resp.StatusCode > http.StatusIMUsed {
		// 200 未満 か 226 超過
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
