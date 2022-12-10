package test

import (
	"testing"
	"user-service/internal/config"
)

func TestGetConfig(t *testing.T) {
	config := config.GetConfig("./../configs/test_config.yml")
	if config.DB.DBname != "test" || config.Redis.DB != 1 || config.DB.Password != "test" {
		t.Fatal("Error: Does not match!")
	}
}
