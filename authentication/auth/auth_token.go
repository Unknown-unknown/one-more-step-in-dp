package auth

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	DefaultExpiredTimeInterval = 1 * 60 // in second
)

type AuthToken struct {
	baseURL         string
	createTimestamp int64
	appID           string
	password        string
	token           string
}

func Create(token string, timestamp int64) *AuthToken {
	return &AuthToken{
		token:           token,
		createTimestamp: timestamp,
	}
}

func CreateAuthToken(baseURL string, createTimestamp int64, appID, password string) *AuthToken {
	return &AuthToken{
		baseURL:         baseURL,
		createTimestamp: createTimestamp,
		appID:           appID,
		password:        password,
	}
}

func (a *AuthToken) Token() string {
	if a.token != "" {
		return a.token
	}
	s := fmt.Sprintf("%s%d%s%s", a.baseURL, a.createTimestamp, a.appID, a.password)
	a.token = fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
	return a.token
}

func (a *AuthToken) IsExpired() bool {
	expireTime := time.Unix(a.createTimestamp, 0).Add(time.Second * DefaultExpiredTimeInterval)
	now := time.Now()
	return now.After(expireTime)
}

func (a *AuthToken) Match(token *AuthToken) bool {
	return a.Token() == token.Token()
}
