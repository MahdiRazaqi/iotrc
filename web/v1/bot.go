package v1

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"

// 	"github.com/labstack/echo"
// )

// func listener(c echo.Context) error {
// 	body, err := ioutil.ReadAll(c.Request().Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var u bot.Update
// 	if err := json.Unmarshal(body, &u); err != nil {
// 		log.Fatalln(err)
// 	}

// 	bot.MessageHandler(*u.Message)

// 	return nil
// }
