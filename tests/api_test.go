package main_test

import (
	"src/api"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	credenciais := api.GetCredentials()

	if credenciais == "" {
		t.Errorf("Error: Something went wrong attempting to get credentials.")
	}
}
