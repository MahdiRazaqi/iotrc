package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/MahdiRazaqi/iotrc/bot"
	"github.com/labstack/echo"
)

func listener(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
	}

	var u bot.Update
	if err := json.Unmarshal(body, &u); err != nil {
		log.Fatalln(err)
	}
	return nil
}
