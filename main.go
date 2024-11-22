package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	router.HandleFunc("/echo", EchoHandler).Methods("POST")
	router.HandleFunc("/time", TimeHandler).Methods("GET")

	fmt.Println(`Server is running on http://localhost:8000

Test commands:
curl http://localhost:8000/health
curl -X POST -H "Content-Type: application/json" -d '{"test": "test"}' http://localhost:8000/echo
curl http://localhost:8000/time`)
	log.Fatal(http.ListenAndServe(":8000", router))
}

// * Test command: curl http://localhost:8000/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server up and running"))
}

// * Test command: curl -X POST -H "Content-Type: application/json" -d '{"test": "test"}' http://localhost:8000/echo
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid json sent in the request payload", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

// * Test command: curl http://localhost:8000/time
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	utcTime := time.Now().UTC()
	timeString := utcTime.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(timeString))
}
