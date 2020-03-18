package auth

import (
	"testing"
)

func TestAuthURL(t *testing.T) {
	au := Init()
	url := "https://caliosd.gq?token=asdfqwert&appID=hjklyuio&timestamp=1584430871"
	au.AuthURL(url)
}
