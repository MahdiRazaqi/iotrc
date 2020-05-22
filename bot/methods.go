package bot

func setWebhook(url string) error {
	_, err := requestToBot("/setWebhook", Params{
		"url": url,
	})
	return err
}

// SendMessage to telegram
func SendMessage(userID int, text string) (Message, error) {
	_, err := requestToBot("/sendMessage", Params{
		"chat_id":    userID,
		"text":       text,
		"parse_mode": "HTML",
	})
	return Message{}, err
}
