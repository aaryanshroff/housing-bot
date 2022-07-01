package main

import (
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
