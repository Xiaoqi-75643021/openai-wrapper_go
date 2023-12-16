package main

import (
	"log"
	"net/http"
	"openai-wrapper/config"
	"openai-wrapper/handlers"
	"os"
	"path/filepath"
)

func main() {
    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        log.Fatal(err)
    }

    logFile, err := os.OpenFile(cfg.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        if os.IsNotExist(err) {
            os.MkdirAll(filepath.Dir(cfg.LogFileName), 0755)
            logFile, err = os.OpenFile(cfg.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
            if err != nil {
                log.Fatalf("Error creating log file: %v", err)
            }
        } else {
            log.Fatalf("Error opening log file: %v", err)
        }
    }
    defer logFile.Close()

    log.SetOutput(logFile)

    // Pass the logger and config to the ChatHandler
    chatHandler := handlers.ChatHandler(cfg, log.New(logFile, "OpenAI-Response: ", log.LstdFlags))

    http.Handle("/chat", chatHandler)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
