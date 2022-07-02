package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestGETWebhook(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/webhook?hub.verify_token=" + os.Getenv("FB_WEBHOOK_VERIFY_TOKEN") + "&hub.challenge=CHALLENGE_ACCEPTED", nil)
    response := httptest.NewRecorder()

    WebhookServer(response, request)

    got := response.Body.String()
    want := "CHALLENGE_ACCEPTED"

    assertResponseBody(t, got, want)
}

func TestPOSTWebhook(t *testing.T) {
    b := []byte(`{"object": "page", "entry": [{"messaging": [{"sender": {"id": "<PSID>"}, "message": {"text": "TEST_MESSAGE"}}]}]}`)

    request, _ := http.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(b));
    response := httptest.NewRecorder()

    WebhookServer(response, request)
    
    var got WebhookResponse
    json.NewDecoder(response.Body).Decode(&got)

    want := WebhookResponse{
        Text: "TEST_MESSAGE",
        PSID: "<PSID>",
    }

    assertResponseBody(t, got, want)
}

func assertResponseBody(t testing.TB, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got %q want %q", got, want)
    }
}
