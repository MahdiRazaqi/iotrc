package bot

func setWebhook(url string) error {
	_, err := requestToBot("/setWebhook", Params{
		"url": url,
	})
	return err
}
