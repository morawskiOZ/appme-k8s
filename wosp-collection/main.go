package main

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	baseValue     = 10000
	hourlyInc     = 1521
	listenAddr    = ":8080"
	expectedUser  = "manhattan"
	expectedPass  = "manhattan"
)

// referenceTime is the epoch from which we calculate hourly increments.
// Set to December 21, 2025, 12:00:00 UTC as the deployment reference.
var referenceTime = time.Date(2025, 12, 21, 12, 0, 0, 0, time.UTC)

type CollectionResponse struct {
	Total int `json:"total"`
}

func main() {
	http.HandleFunc("/pm2go/test/basicauth", handleCollection)
	http.HandleFunc("/health", handleHealth)

	log.Printf("Starting wosp-collection server on %s", listenAddr)
	log.Printf("Reference time: %s", referenceTime.Format(time.RFC3339))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func handleCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !checkBasicAuth(r) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	total := calculateTotal()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CollectionResponse{Total: total})
}

func checkBasicAuth(r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	userMatch := subtle.ConstantTimeCompare([]byte(user), []byte(expectedUser)) == 1
	passMatch := subtle.ConstantTimeCompare([]byte(pass), []byte(expectedPass)) == 1

	return userMatch && passMatch
}

func calculateTotal() int {
	now := time.Now().UTC()
	elapsed := now.Sub(referenceTime)

	if elapsed < 0 {
		return baseValue
	}

	hours := int(elapsed.Hours())
	return baseValue + (hours * hourlyInc)
}
