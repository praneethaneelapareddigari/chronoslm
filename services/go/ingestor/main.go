package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status":"ok","service":"ingestor"}`)
	})
	http.HandleFunc("/simulate-ingest", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		fmt.Fprintln(w, `{"ingested": 10000, "duration_ms": 20}`)
	})
	log.Println("Ingestor service listening on :7070")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
