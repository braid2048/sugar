package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SendGetRequest 发送GET请求
func SendGetRequest(ctx context.Context, url string, headers map[string]string) (int, []byte, error) {
	client := http.DefaultClient

	// 创建一个GET请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("init req err: %w", err)
	}

	// 设置header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("do req err: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("read res err: %w", err)
	}

	return resp.StatusCode, body, nil
}

// SendPostRequest 发送POST请求
func SendPostRequest(ctx context.Context, url string, headers map[string]string, body []byte) (int, []byte, error) {
	client := http.DefaultClient

	// 创建一个POST请求
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, fmt.Errorf("init req err: %w", err)
	}

	// 设置header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("do req err: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("read res err: %w", err)
	}

	return resp.StatusCode, respBody, nil
}

// SendPOSTFormData 发送POST请求表单参数
func SendPOSTFormData(ctx context.Context, reqURL string, headers map[string]string, params map[string]string) (int, []byte, error) {
	client := http.DefaultClient
	formData := url.Values{}

	for key, value := range params {
		formData.Set(key, value)
	}
	// 发起 POST 请求
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return 0, nil, fmt.Errorf("SendPOSTFormData -- init req err: %w", err)
	}
	// 设置header
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("do req err: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("read res err: %w", err)
	}

	return resp.StatusCode, respBody, nil
}
