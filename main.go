package main

import (
    "log"
    "net/http"
)

func main() {
    handler := http.HandlerFunc(WebhookServer)
    log.Fatal(http.ListenAndServe(":80", handler))
}
