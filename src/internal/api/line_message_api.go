package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"football_tracker/src/internal/config"
	"football_tracker/src/internal/models"
)

type LineMessageClient struct {
	channelAccessToken string
	apiURL             string
	httpClient         *http.Client
}

func NewLineClient() *LineMessageClient {
	return &LineMessageClient{
		channelAccessToken: config.LineChannelAccessToken,
		apiURL:             config.LineAPIURL,
		httpClient:         &http.Client{},
	}
}

func (c *LineMessageClient) PushMessage(message models.LinePayload) error {

	jsonData, err := json.Marshal(message)

	if err != nil {
		return fmt.Errorf("error marshalling push message: %w", err)
	}

	req, err := http.NewRequest("POST", c.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.channelAccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		log.Printf("LINE API request failed with status code: %d, response: %s", resp.StatusCode, buf.String())
		return fmt.Errorf("LINE API request failed with status code: %d", resp.StatusCode)
	}

	return nil
}
