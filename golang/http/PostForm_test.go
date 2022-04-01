package http_test

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func ExamplePostForm() {
	var (
		requestURL = `https://github.com/robfig/cron/issues`
		outputFile = "resp.html"
	)

	resp, err := http.PostForm(requestURL, url.Values{
		"is": []string{"issue", "open"},
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outputFile, body, 0666)
	if err != nil {
		panic(err)
	}
	// Output:
}
