package service

import (
	"os"
	"strings"
)

type AuthConfig struct {
	Username string
	Password string
}

type IAMConfig struct {
	Region            string
	Environment       string
	OrgID             string
	ServiceID         string
	ServicePrivateKey string
}

// SecretFromEnv returns the value of the given string, if it is an environment variable.
func SecretFromEnv(value string) string {
	if strings.HasPrefix(value, "$") {
		envVar := strings.TrimPrefix(value, "$")
		if v := os.Getenv(envVar); v != "" {
			return v
		}
	}
	return value
}
