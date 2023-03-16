package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type application struct {
}

const signatureKeyConst = "signature-key"
const apiKeyKeyConst = "x-api-key"
const apiKeyValueConst = "sample-key"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/go/nethttp/api-key-signature-protected", apiKeyAuth(signatureAuth(webhookPluginAPI)))

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

// Exposed protected handler
func webhookPluginAPI(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "200 OK")
	w.WriteHeader(http.StatusOK)
	return
}

// Verify api key and timestamp
func apiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Match api key from request header
		apiKeyMatch := r.Header.Get(apiKeyKeyConst) == apiKeyValueConst

		if !apiKeyMatch {
			logRequest(r, "Api key auth Failed")
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		}

		// Check request timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
		requestTimestamp, err := strconv.ParseInt(r.Header.Get("conceal_timestamp"), 10, 64)
		currentTimestamp := time.Now().Unix()
		if err != nil {
			logRequest(r, "Invalid Timestamp")
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Invalid Timestamp", http.StatusBadRequest)
		}
		log.Println(fmt.Sprintf("Time Diff: %d", requestTimestamp-currentTimestamp))
		if requestTimestamp-currentTimestamp < -60000 || requestTimestamp-currentTimestamp > 120000 {
			logRequest(r, "Invalid Timestamp. Timestamp out of range")
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Invalid Timestamp. Timestamp out of range", http.StatusBadRequest)
		}

		next.ServeHTTP(w, r)
		return
	})
}

// Verify signature value
func signatureAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timestamp := r.Header.Get("conceal_timestamp")
		messageSignature := r.Header.Get("conceal_signature")

		message := fmt.Sprintf("%s|%s://%s%s", timestamp, "http", r.Host, r.URL.Path)
		hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
		hasher.Write([]byte(message))
		sha := fmt.Sprintf("%x", hasher.Sum(nil))

		signatureMatch := sha == messageSignature

		if !signatureMatch {
			logRequest(r, "Signature auth Failed")
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Invalid Signature", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
		return
	})
}

func logRequest(r *http.Request, tag string) {
	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	log.Println(fmt.Sprintf("Request Status: %s", tag))
	log.Println("REQUEST")
	log.Println(fmt.Sprintf("Url : %s?%s", r.URL, r.URL.RawQuery))
	log.Println(fmt.Sprintf("Method : %s", r.Method))
	log.Println(fmt.Sprintf("Header : %s", r.Header))
	log.Println(fmt.Sprintf("Body: %s", body))
}
