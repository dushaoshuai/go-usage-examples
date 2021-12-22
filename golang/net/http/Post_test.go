package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type request struct {
	Source     string `json:"source"`
	UserID     uint64 `json:"user_id,string"`
	SID        string `json:"sid"`
	OperatorID uint64 `json:"operator_id,string"`
	Data       struct {
		OrderID uint64 `json:"order_id,string"`
		Content struct {
			Text   string   `json:"text"`
			Images []string `json:"images"`
		} `json:"content"`
		OrderType int `json:"order_type"`
	} `json:"data"`
}

func Example_post() {
	url := "http://dev2.sponsor.test.collectivedynamic.com:30000/sponsor/leave_message/create"

	data, err := json.Marshal(request{
		Source:     "web",
		UserID:     1184298124261195776,
		SID:        "ZjgxMzg4MzBhZGRkMjZkZjgxZmNhYjJlOGUyYmQ0ZDcxYTQ5OGYwOV8xNDMxMTkzNzI1ODIzODc3MTIw",
		OperatorID: 1431193725823877120,
		Data: struct {
			OrderID uint64 `json:"order_id,string"`
			Content struct {
				Text   string   `json:"text"`
				Images []string `json:"images"`
			} `json:"content"`
			OrderType int `json:"order_type"`
		}{
			OrderID: 1473525462092218368,
			Content: struct {
				Text   string   `json:"text"`
				Images []string `json:"images"`
			}{
				Text:   "三狗子家的也行",
				Images: nil,
			},
			OrderType: 4,
		},
	})
	if err != nil {
		panic(err)
	}

	payload := strings.NewReader(string(data))

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = rsp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	if _, err = os.Stdout.Write(body); err != nil {
		panic(err)
	}
	// Output:
	// {
	//  "error_code": 0,
	//  "error_description": "",
	//  "data": {}
	// }
}
