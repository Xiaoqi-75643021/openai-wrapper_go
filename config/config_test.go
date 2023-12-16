package config

import (
    "os"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    // Set an environment variable to test the override feature
    os.Setenv("OPENAI_API_KEY", "test_api_key")
    defer os.Unsetenv("OPENAI_API_KEY") // Clean up the environment variable after the test

    // Assume a valid configuration file path
    validConfigPath := "../config/config.json"

    // Test successful configuration loading
    cfg, err := LoadConfig(validConfigPath)
    if err != nil {
        t.Fatalf("LoadConfig() error = %v, wantErr %v", err, false)
    }

    // Ensure the environment variable has overridden the key from the file
    if cfg.OpenAIKey != "test_api_key" {
        t.Errorf("LoadConfig() did not override OpenAIKey with environment variable")
    }

    // Assume an invalid configuration file path
    invalidConfigPath := "../config/nonexistent.json"

    // Attempt to load a non-existent configuration file
    _, err = LoadConfig(invalidConfigPath)
    if err == nil {
        t.Errorf("LoadConfig() error = %v, wantErr %v", err, true)
    }
}
