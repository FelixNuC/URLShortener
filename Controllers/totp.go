package controllers

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(secret string) (string, error) {
	totpCode, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return totpCode, nil
}

func VerifyTOTP(totpCode, secret string) bool {
	return totp.Validate(totpCode, secret)
}
