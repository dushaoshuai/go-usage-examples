package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type shortURL struct {
	Sid      string `json:"sid"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
	EnterURL string `json:"enter_url"`
	Name     string `json:"name"`
}

type shortURLResp struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	ShortURL shortURL `json:"short_url"`
}

func Example_yunpian_short_url() {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	form := url.Values{}
	form.Add("apikey", "api-xxxxxxxxxxxxxxx-key")
	form.Add("long_url", "https://github.com/gin-gonic")
	reqBody := strings.NewReader(form.Encode())

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost,
		"https://sms.yunpian.com/v2/short_url/shorten.json", reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept", "charset=utf-8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "charset=utf-8")
	req.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var shortURLResp shortURLResp
	err = json.Unmarshal(respBody, &shortURLResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", shortURLResp)
}
