package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	URL = "https://apply.stanford.edu/portal/campus-visit?cmd=getDates&dtstart=2022-06-17&dtend=2022-06-21"
)

type CalendarResponse struct {
	Dates [][]string `json:"dates"`
}

func HandleRequest(ctx context.Context) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	found := CalendarResponse{}
	err = json.Unmarshal(html, &found)
	if err != nil {
		return err
	}

	for _, value := range found.Dates {
		full, _ := strconv.Atoi(value[1])
		if full == 0 {
			err := SendWebhook(value[0])
			if err != nil {
				return err
			}
		}
		fmt.Println(value)
	}

	return nil
}

type WebhookRequest struct {
	Content string `json:"content"`
}

func SendWebhook(day string) error {
	url := os.Getenv("WEBHOOK_URL")
	if len(url) == 0 {
		panic("Cannot load environment variable `WEBHOOK_URL`")
	}
	payload, err := json.Marshal(WebhookRequest{Content: day + " is available. @everyone"})
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(payload))
	return err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	lambda.Start(HandleRequest)
}
