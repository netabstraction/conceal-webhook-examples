package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	signatureKeyConst = "signature-key"
	apiKeyKeyConst    = "x-api-key"
	apiKeyValueConst  = "sample-key"
	webhookAddress    = "http://127.0.0.1:8080/webhook"
)

func main() {

	// Create a new HTTP server
	http.HandleFunc("/webhook", handleWebhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// Exposed Webhook
func handleWebhook(w http.ResponseWriter, r *http.Request) {

	requestApiKey := r.Header.Get(apiKeyKeyConst)
	requestTimestamp := r.Header.Get("conceal-timestamp")
	requestSignature := r.Header.Get("conceal-signature")

	//Api Key validation
	if requestApiKey != apiKeyValueConst {
		log.Printf("Invalid API Key")
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	//Timestamp validation
	if !isValidTimestamp(requestTimestamp) {
		log.Printf("Invalid Timestamp")
		http.Error(w, "Invalid Timestamp", http.StatusBadRequest)
		return
	}

	//Signature validation
	if !isValidSignature(requestTimestamp, requestSignature) {
		log.Printf("Invalid Signature")
		http.Error(w, "Invalid Signature", http.StatusUnauthorized)
		return
	}

	// Process the webhook payload
	// ..
	logRequest(w, r)
	// ..

	// Return a success response
	log.Printf("OK")
	w.WriteHeader(http.StatusOK)
	return
}

// Validate timestamp timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
func isValidTimestamp(requestTimestamp string) bool {
	requestTimestampInt, err := strconv.ParseInt(requestTimestamp, 10, 64)
	currentTimestamp := time.Now().Unix()
	if err != nil {
		return false
	}
	log.Printf(fmt.Sprintf("Time Diff: %d", requestTimestampInt-currentTimestamp))
	if requestTimestampInt-currentTimestamp < -60000 || requestTimestampInt-currentTimestamp > 120000 {
		return false
	}

	return true
}

// Validate signature
func isValidSignature(requestTimestamp string, requestSignature string) bool {

	message := fmt.Sprintf("%s|%s", requestTimestamp, webhookAddress)
	hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
	hasher.Write([]byte(message))
	sha := fmt.Sprintf("%x", hasher.Sum(nil))

	return sha == requestSignature
}

// Log request
func logRequest(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Invalid Body")
		http.Error(w, "Invalid Body", http.StatusUnauthorized)
		return
	}
	log.Printf(fmt.Sprintf("req [%s] %s", r.Method, r.URL))
	log.Printf(fmt.Sprintf("headers : %s", r.Header))
	log.Printf(fmt.Sprintf("body: %v", body))
}
