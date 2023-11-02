package service_test

import (
	"github.com/loafoe/iamsale/service"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

// Write tests for the service package here
func TestSecretFromEnv(t *testing.T) {
	t.Parallel()

	os.Setenv("TEST_ENV_VAR", "test")
	cfg := service.IAMConfig{
		ServicePrivateKey: "$TEST_ENV_VAR",
	}
	val := service.SecretFromEnv(cfg.ServicePrivateKey)
	assert.Equal(t, val, "test")
}
