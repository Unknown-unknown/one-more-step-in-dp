package auth

import (
	"log"
	"net/url"
	"strconv"
)

type APIRequest struct {
	baseURL   string
	token     string
	appID     string
	timestamp int64
}

func CreateFromFullURL(fullURL string) *APIRequest {
	u, err := url.Parse(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	t, err := strconv.Atoi(q.Get("timestamp"))
	if err != nil {
		log.Fatal(err)
	}
	return &APIRequest{
		baseURL:   u.Host,
		token:     q.Get("token"),
		appID:     q.Get("appID"),
		timestamp: int64(t),
	}
}

func CreateAPIRequest(baseURL, token, appID string, timestamp int64) *APIRequest {
	return &APIRequest{
		baseURL:   baseURL,
		token:     token,
		appID:     appID,
		timestamp: timestamp,
	}
}

func (r *APIRequest) BaseURL() string {
	return r.baseURL
}

func (r *APIRequest) Token() string {
	return r.token
}

func (r *APIRequest) APPID() string {
	return r.appID
}

func (r *APIRequest) Timestamp() int64 {
	return r.timestamp
}
