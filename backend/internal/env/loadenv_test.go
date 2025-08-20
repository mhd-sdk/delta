package env

import (
	"os"
	"testing"
)

func TestLoadEnv_MissingVariables(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("API_KEY")
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("MARKET_DATA_URL")
	os.Unsetenv("ROUTING_URL")

	err := LoadEnv()
	if err == nil {
		t.Error("Expected error for missing environment variables")
	}

	if err.Error() != ErrLoadingEnv.Error() {
		t.Errorf("Expected ErrMissingEnvVars, got %v", err)
	}
}

func TestLoadEnv_Success(t *testing.T) {
	// Set required environment variables
	os.Setenv("API_KEY", "test-api-key")
	os.Setenv("SECRET_KEY", "test-secret-key")
	os.Setenv("MARKET_DATA_URL", "https://test-market.com")
	os.Setenv("ROUTING_URL", "https://test-routing.com")

	// Create a temporary .env file for testing
	envContent := `API_KEY=test-api-key
SECRET_KEY=test-secret-key
MARKET_DATA_URL=https://test-market.com
ROUTING_URL=https://test-routing.com`

	err := os.WriteFile(".env", []byte(envContent), 0644)
	if err != nil {
		t.Skipf("Could not create .env file: %v", err)
	}
	defer os.Remove(".env")

	err = LoadEnv()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Clean up
	os.Unsetenv("API_KEY")
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("MARKET_DATA_URL")
	os.Unsetenv("ROUTING_URL")
}

func TestErrorConstants(t *testing.T) {
	if ErrMissingEnvVars.Error() != "missing environment variables" {
		t.Errorf("Expected 'missing environment variables', got '%s'", ErrMissingEnvVars.Error())
	}

	if ErrLoadingEnv.Error() != "could not load environment variables" {
		t.Errorf("Expected 'could not load environment variables', got '%s'", ErrLoadingEnv.Error())
	}
}
