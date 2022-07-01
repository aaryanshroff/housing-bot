package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
)

type Messaging struct {
    Message string `json:"message"`
}

type Entry struct {
    Messaging []Messaging `json:"messaging"`
}

type Body struct {
    Object string  `json:"object"`
    Entry  []Entry `json:"entry"`
}

func main () {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my Go bot!")
    })

    http.HandleFunc("/webhook", WebhookHandler)
    
    http.ListenAndServe(":80", nil)
}

func WebhookHandler (w http.ResponseWriter, r *http.Request) {
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
        var b Body
        json.NewDecoder(r.Body).Decode(&b)
        if b.Object == "page" {
            fmt.Fprintf(w, b.Entry[0].Messaging[0].Message)  
        }
    }
}
