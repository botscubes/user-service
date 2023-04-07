package test

import (
	"testing"

	"github.com/botscubes/user-service/internal/config"
)

func TestGetConfig(t *testing.T) {
	config, err := config.GetConfig("./../configs/test_config.yml")
	if err != nil && (config.DB.DBname != "test" || config.Redis.DB != 1 || config.DB.Password != "test") {
		t.Fatalf("Error: %v", err)
	}
}
