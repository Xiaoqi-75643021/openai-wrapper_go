package openai

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
)

// Create a test HTTP client function for mocking OpenAI API responses
type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req), nil
}

func TestClient_Chat(t *testing.T) {
    // Create a mock HTTP client
    client := &Client{
        APIKey: "test_api_key",
        APIURL: "http://mock.api/",
    }

    // Set up the mock response
    client.httpClient = &http.Client{
        Transport: RoundTripFunc(func(req *http.Request) *http.Response {
            return &http.Response{
                StatusCode: http.StatusOK,
                Body:       io.NopCloser(bytes.NewBufferString(`{"id":"123", "choices": [{"message": {"content": "test response"}}]}`)),
                Header:     make(http.Header),
            }
        }),
    }

    // Construct a request body
    chatRequest := &ChatRequest{
        Model: "gpt-3.5-turbo",
        Messages: []Message{
            {
                Role:    "user",
                Content: "Hello, who are you?",
            },
        },
    }

	logger := log.New(io.Discard, "", log.LstdFlags)

    // Call the Chat method
    resp, err := client.Chat(chatRequest, logger)
    if err != nil {
        t.Errorf("Client.Chat() error = %v", err)
    }

    // Check if the response is as expected
    expectedResponseContent := "test response"
    if resp.Choices[0].Message.Content != expectedResponseContent {
        t.Errorf("Client.Chat() got = %v, want %v", resp.Choices[0].Message.Content, expectedResponseContent)
    }
}
