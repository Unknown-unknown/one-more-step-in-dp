package test

import (
	"one-more-step-in-dp/authentication/auth"
	"testing"
)

func TestAuthURL(t *testing.T) {
	au := auth.Init()
	url := "https://caliosd.gq?token=asdfqwert&appID=hjklyuio&timestamp=1584430871"
	au.AuthURL(url)
}
