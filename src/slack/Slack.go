package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Comment struct {
	Message  string `json:"text"`
	Username string `json:"username"`
	IconUrl  string `json:"icon_url"`
}

type Config struct {
	Webhook string
}

func (config Config) Post(c Comment) error {
	json, err := json.Marshal(c)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.Webhook, bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}

func DefaultComment(message string) Comment {
	return Comment{Message: message, Username: "Chucks Tap Tracer", IconUrl: "https://emoji.slack-edge.com/T02EA9PT09H/chucks/d131fc4b010c32db.png"}
}
