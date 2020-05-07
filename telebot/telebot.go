package telebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var endpoint string = "https://api.telegram.org/bot1296532223:AAFhySQkXvZCkQbKANGgaRCa8TU3KBq4etk/"

// Start bot server
func Start() {
	var offset int
	for {
		resp := getUpdates(Params{
			"offset": offset + 1,
		})

		for _, m := range resp.Result {
			offset = m.UpdateID
			controller(m.Message)
		}
	}
}

func request(query string, method string, params Params) TelegramResponse {
	var url string
	url = endpoint + query

	if method == http.MethodGet {
		url += "?"
		for key, value := range params {
			url += fmt.Sprintf("%s=%v&", key, value)
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []byte(body)
	var response TelegramResponse
	if err := json.Unmarshal(jsonData, &response); err != nil {
		log.Println(err)
	}
	return response
}
