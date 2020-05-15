package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var endpoint string = "https://api.telegram.org/bot1296532223:AAFhySQkXvZCkQbKANGgaRCa8TU3KBq4etk"

// Start bot server
func Start() {
	if err := setWebhook("https://test.com/bot"); err != nil {
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

func requestToBot(path string, params Params) (interface{}, error) {
	resp, err := request(endpoint+path, http.MethodGet, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	if !response["ok"].(bool) {
		return nil, errors.New("unsuccessful request")
	}

	return response, nil
}
