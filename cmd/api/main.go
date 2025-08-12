package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoapi/internal/domain/model/dto"
)

func main() {
	//Simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte("OK"))
		// data := map[string]string{
		// 	"status": "healthy",
		// }
		data := dto.Todo{}

		Resp, err := json.Marshal(data)
		if err != nil {
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(Resp)
	})
	//http.ListenAndServe(":8080", nil)
	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
