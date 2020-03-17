package auth

type CredentialStorage interface {
	GetPasswordByAPPID(appID string) string
}

type MySQLCredentialStorage struct{}

func (s *MySQLCredentialStorage) GetPasswordByAPPID(appID string) string {
	return ""
}
