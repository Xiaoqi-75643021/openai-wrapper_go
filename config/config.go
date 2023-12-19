package config

import (
    "encoding/json"
    "os"
)

type Configuration struct {
    OpenAIKey   string            `json:"OpenAIKey"`
    OpenAIURL   string            `json:"OpenAIURL"`
    LogFileName string            `json:"LogFileName"`
    ModelsList  map[string]bool   `json:"ModelsList"`
}

func LoadConfig(path string) (Configuration, error) {
    var config Configuration
    configFile, err := os.Open(path)
    if err != nil {
        return Configuration{}, err
    }
    defer configFile.Close()

    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        return Configuration{}, err
    }

    OverrideConfigWithEnv(&config)
    
    return config, err
}

// Override sensitive information(OpenAI_apiKey) in the configuration with environment variables
func OverrideConfigWithEnv(cfg *Configuration) {
    if apiKey := os.Getenv("OPENAI_API_KEY"); apiKey != "" {
        cfg.OpenAIKey = apiKey
    }
}