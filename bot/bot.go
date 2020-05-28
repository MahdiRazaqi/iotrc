package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MahdiRazaqi/iotrc/sensorlog"
	"go.mongodb.org/mongo-driver/bson"
)

var endpoint string = "https://api.telegram.org/bot1296532223:AAFhySQkXvZCkQbKANGgaRCa8TU3KBq4etk"

// Start bot server
func Start() {
	if err := setWebhook("https://5ecf7d78bc51.ngrok.io/api/v1/bot"); err != nil {
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

// MessageHandler for handle commands
func MessageHandler(message Message) {
	switch message.Text {
	case "/start":
		SendMessage(message.From.ID, "Welcome to IOTRC!")
		break
	case "/get_logs":
		sensorlogs, err := sensorlog.Find(bson.M{}, 1, 10)
		if err != nil {
			log.Fatalln(err)
			return
		}

		var msg string
		for _, s := range sensorlogs {
			if msg != "" {
				msg += fmt.Sprintf("%%0A%%0A%%0A")
			}
			msg += fmt.Sprintf("<b>Temp: </b>%.2f%%0A<b>Humidity: </b>%.2f%%0A<b>DustHumidity: </b>%.2f%%0A<b>Light: </b>%.2f%%0A<b>Pomp Status: </b>%v%%0A<b>Lamp Status: </b>%v%%0A<code>%v</code>", s.TempDHT, s.HumidityDHT, s.DustHumidity, s.Light, s.PompStatus, s.LampStatus, s.Created.Format("02-01-2006 15:04:05"))
		}

		SendMessage(message.From.ID, msg)
		break
	case "/turn_on_pomp":
		SendMessage(message.From.ID, "turn_on_pomp")
		break
	case "/turn_off_pomp":
		SendMessage(message.From.ID, "turn_off_pomp")
		break
	case "/turn_on_lamp":
		SendMessage(message.From.ID, "turn_on_lamp")
		break
	case "/turn_off_lamp":
		SendMessage(message.From.ID, "turn_off_lamp")
		break
	}
}
