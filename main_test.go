package main

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestGETWebhook(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/webhook?hub.verify_token=" + os.Getenv("FB_WEBHOOK_VERIFY_TOKEN") + "&hub.challenge=CHALLENGE_ACCEPTED", nil)
    response := httptest.NewRecorder()

    WebhookHandler(response, request)

    got := response.Body.String()
    want := "CHALLENGE_ACCEPTED"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func TestPOSTWebhook(t *testing.T) {
    b := []byte(`{"object": "page", "entry": [{"messaging": [{"message": "TEST_MESSAGE"}]}]}`)

    request, _ := http.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(b));
    response := httptest.NewRecorder()

    WebhookHandler(response, request)

    got := response.Body.String()
    want := "TEST_MESSAGE"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
