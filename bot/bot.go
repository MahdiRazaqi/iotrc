package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var endpoint string = "https://api.telegram.org/bot1296532223:AAFhySQkXvZCkQbKANGgaRCa8TU3KBq4etk"

// Start bot server
func Start() {
	if err := setWebhook("https://6ad8ff38.ngrok.io/api/v1/bot"); err != nil {
		log.Fatalln(err)
	}
}

func request(url, method string, params Params) (*http.Response, error) {
	var body io.Reader

	if method == http.MethodGet {
		if len(params) != 0 {
			url += "?"
			for key, value := range params {
				url += fmt.Sprintf("%s=%v&", key, value)
			}
		}
	} else {
		jsonData, _ := json.Marshal(params)
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func proxy(path string, params Params) (*http.Response, error) {
	var contentData string
	for k, v := range params {
		if contentData != "" {
			contentData += "&"
		}
		contentData += k + "=" + fmt.Sprintf("%v", v)
	}

	json, err := json.Marshal(map[string]interface{}{
		"UrlBox":     endpoint + path + "?" + contentData,
		"MethodList": http.MethodGet,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://www.httpdebugger.com/tools/ViewHttpHeaders.aspx", "application/json", bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
