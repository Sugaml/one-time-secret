package util

import "github.com/google/uuid"

func GetToken() string {
	token := uuid.New()
	return token.String()
}
