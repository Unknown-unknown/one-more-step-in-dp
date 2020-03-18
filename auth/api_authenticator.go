package auth

import "log"

type APIAuthenticator interface {
	AuthURL(url string)
	AuthReq(req *APIRequest)
}

type DefaultAuthenticator struct {
	credentialStorage CredentialStorage
}

func Init() *DefaultAuthenticator {
	return &DefaultAuthenticator{
		credentialStorage: &MySQLCredentialStorage{},
	}
}

func (da *DefaultAuthenticator) AuthURL(url string) {
	req := CreateFromFullURL(url)
	da.AuthReq(req)
}

func (da *DefaultAuthenticator) AuthReq(req *APIRequest) {
	appID, token := req.APPID(), req.Token()
	timestamp := req.Timestamp()
	baseURL := req.BaseURL()

	clientAuthToken := Create(token, timestamp)
	if clientAuthToken.IsExpired() {
		log.Fatalln("Token is expired.")
	}

	password := da.credentialStorage.GetPasswordByAPPID(appID)
	serverAuthToken := CreateAuthToken(baseURL, timestamp, appID, password)
	if !serverAuthToken.Match(clientAuthToken) {
		log.Fatalln("Token verification failed.")
	}
}
