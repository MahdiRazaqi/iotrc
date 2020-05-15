package bot

import "fmt"

func setWebhook(url string) error {
	ee, err := requestToBot("/setWebhook", Params{
		"url": url,
	})
	fmt.Println(ee)
	return err
}
