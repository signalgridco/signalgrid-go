package signalgrid

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	ClientKey string
	Endpoint  string
	HTTP      *http.Client
}

func NewClient(clientKey string) (*Client, error) {
	if clientKey == "" {
		return nil, errors.New("client key is required")
	}

	return &Client{
		ClientKey: clientKey,
		Endpoint:  "https://api.signalgrid.co",
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

type Message struct {
	Channel  string
	Type     string
	Title    string
	Body     string
	Critical bool
}

func (c *Client) Send(msg Message) (string, error) {
	if msg.Channel == "" {
		return "", errors.New("channel token is required")
	}

	data := url.Values{}
	data.Set("client_key", c.ClientKey)
	data.Set("channel", msg.Channel)

	if msg.Type != "" {
		data.Set("type", msg.Type)
	}
	if msg.Title != "" {
		data.Set("title", msg.Title)
	}
	if msg.Body != "" {
		data.Set("body", msg.Body)
	}
	if msg.Critical {
		data.Set("critical", "true")
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.Endpoint,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", errors.New(string(body))
	}

	return string(body), nil
}