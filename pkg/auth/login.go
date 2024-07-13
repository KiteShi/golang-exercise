package auth

import (
	"errors"
	"strings"
)

var admins = make(map[string]string)

var errUninitializedCreds = errors.New("empty administrator credentials")

func InitAdmins(rawCreds string) error {
	if len(rawCreds) == 0 {
		return errUninitializedCreds
	}

	adminPairs := strings.Split(rawCreds, ",")
	for _, pair := range adminPairs {
		credentials := strings.Split(pair, ":")
		if len(credentials) == 2 {
			admins[credentials[0]] = credentials[1]
		}
	}

	if err := validateAdminsCredentials(); err != nil {
		return err
	}

	return nil
}

func Authenticate(username, password string) (bool, error) {
	if err := validateAdminsCredentials(); err != nil {
		return false, err
	}

	if pwd, exists := admins[username]; !exists || pwd != password {
		return false, nil
	}

	return true, nil
}

func validateAdminsCredentials() error {
	if len(admins) == 0 {
		return errUninitializedCreds
	}

	return nil
}
