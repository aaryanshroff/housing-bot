package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
)

type Sender struct {
    ID string `json:"id"`
}

type Message struct {
    Text string `json:"text"`
}

type Messaging struct {
    Sender  Sender  `json:"sender"`
    Message Message `json:"message"`
}

type Entry struct {
    Messaging []Messaging `json:"messaging"`
}

type WebhookRequestBody struct {
    Object string  `json:"object"`
    Entry  []Entry `json:"entry"`
}

type WebhookResponse struct {
    Text string `json:"text"`
    PSID string `json:"psid"`
}

func WebhookServer (w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        c := r.URL.Query().Get("hub.challenge")
        t := r.URL.Query().Get("hub.verify_token")

        if t == os.Getenv("FB_WEBHOOK_VERIFY_TOKEN") {
            fmt.Fprintf(w, c)
        } else {
            w.WriteHeader(http.StatusUnauthorized)
        }

    case http.MethodPost:
        var b WebhookRequestBody
        json.NewDecoder(r.Body).Decode(&b)
        if b.Object == "page" {
            text := b.Entry[0].Messaging[0].Message.Text
            psid := b.Entry[0].Messaging[0].Sender.ID
            res := WebhookResponse{
                Text: text,
                PSID: psid,
            }
            json.NewEncoder(w).Encode(res)
        }
    }
}
