package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "openai-wrapper/config"
    "openai-wrapper/openai"
)

// ChatHandler handles the /chat endpoint
func ChatHandler(cfg config.Configuration, logger *log.Logger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }

        var chatRequest openai.ChatRequest
        if err := json.NewDecoder(r.Body).Decode(&chatRequest); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer r.Body.Close() // Ensure the request body is closed

        if chatRequest.Model == "" || len(chatRequest.Messages) == 0 {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
        }

        // Check if the requested model is in the list of available models
        if _, exists := cfg.ModelsList[chatRequest.Model]; !exists {
            http.Error(w, "Model not supported", http.StatusBadRequest)
            return
        }

        client := openai.NewClient(cfg.OpenAIKey, cfg.OpenAIURL)
        chatResponse, err := client.Chat(&chatRequest, logger) // Pass the logger to the Chat function
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(chatResponse)
    }
}
