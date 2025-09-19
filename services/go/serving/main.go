package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Hit struct {
	ID    string  `json:"id"`
	Score float64 `json:"score"`
}

type Resp struct {
	Hits      []Hit   `json:"hits"`
	LatencyMS int64   `json:"latency_ms"`
	P95       float64 `json:"p95,omitempty"`
	P99       float64 `json:"p99,omitempty"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(time.Duration(10+rand.Intn(40)) * time.Millisecond)
	resp := Resp{
		Hits: []Hit{{ID: "doc_123", Score: 0.97}, {ID: "doc_880", Score: 0.94}},
		LatencyMS: time.Since(start).Milliseconds(),
		P95:       120,
		P99:       148,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok","service":"serving"}`))
	})
	http.HandleFunc("/search", searchHandler)
	log.Println("Serving service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
