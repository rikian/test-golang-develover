package config

import "testing"

func TestConfig(t *testing.T) {
	LoadEnvFile()

	if err := ConnectDB().Error; err != nil {
		t.Fatal(err.Error())
	}
}
