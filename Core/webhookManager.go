package Core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func WebhookSend() {
	WebhookURL := "https://discord.com/api/webhooks/"

	payloadData := struct {
		Content   interface{} `json:"content"`
		Embeds    []Embed     `json:"embeds"`
		Username  string      `json:"username"`
		AvatarURL string      `json:"avatar_url"`
	}{
		Content: nil,
		Embeds: []Embed{
			{
				Title:       "NEW ELECEED CHAPTER FOUND",
				Description: "The latest chapter has been found",
				Color:       5814783,
				Fields: []Field{
					{
						Name:  "LINK BELOW",
						Value: "placeholder",
					},
				},
			},
		},
		Username:  "Monitor",
		AvatarURL: "https://i.imgur.com/gTtPuMp.png",
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		log.Fatal("Encoding json failed")
	}

	req, err := http.NewRequest("POST", WebhookURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("couldn't create webhook")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Coudln't send request")
	}

	defer resp.Body.Close()
}
