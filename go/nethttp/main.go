package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

const basicAuthUserNameConst = "alice"
const basicAuthPasswordConst = "p8fnxeqj5a7zbrqp"
const signatureKeyConst = "signature-key"
const apiKeyKeyConst = "x-api-key"
const apiKeyValueConst = "sample-key"


func main() {
	app := new(application)

	mux := http.NewServeMux()
	mux.HandleFunc("/unprotected", app.unprotectedHandler)
	mux.HandleFunc("/api-key-protected", app.apiKeyAuth(app.protectedHandler))
	mux.HandleFunc("/api-key-signature-protected", app.apiKeyAuth(app.signatureAuth(app.protectedHandler)))

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

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r, "Protected Handler")
	fmt.Fprintln(w, "This is the protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r, "Unprotected handler")
	fmt.Fprintln(w, "This is the unprotected handler")
}

func (app *application) apiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKeyMatch := r.Header.Get(apiKeyKeyConst) == apiKeyValueConst

		if apiKeyMatch {
			next.ServeHTTP(w, r)
			return
		}

		app.logRequest(r, "Api key auth Failed")

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	})
}

func (app *application) signatureAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timestamp := r.Header.Get("conceal_timestamp")
		messageSignature := r.Header.Get("conceal_signature")

		message := fmt.Sprintf("%s|%s://%s%s", timestamp, "http", r.Host, r.URL.Path)
		hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
		hasher.Write([]byte(message))
		sha := fmt.Sprintf("%x", hasher.Sum(nil))

		signatureMatch := sha == messageSignature

		if signatureMatch {
			next.ServeHTTP(w, r)
			return
		}

		app.logRequest(r, "Signature auth Failed")

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Invalid Signature", http.StatusUnauthorized)
	})
}

func (app *application) logRequest(r *http.Request, tag string) {
	log.Println("Got a new request " + tag)
	log.Println(r.URL)
	log.Println(r.Method)
	log.Println(r.Header)
	log.Println(r.URL.RawQuery)

	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	log.Println(body)
}
