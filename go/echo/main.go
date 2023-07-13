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

	"github.com/labstack/echo/v4"
)

const (
	signatureKeyConst = "signature-key"
	apiKeyKeyConst    = "x-api-key"
	apiKeyValueConst  = "sample-key"
	webhookAddress    = "http://127.0.0.1:8080/webhook"
)

type Error struct {
	Error string `json:"error"`
}

type Status struct {
	Status string `json:"status"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.POST("/webhook", handleWebhook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Exposed Webhook
func handleWebhook(c echo.Context) error {

	requestApiKey := c.Request().Header.Get(apiKeyKeyConst)
	requestTimestamp := c.Request().Header.Get("conceal-timestamp")
	requestSignature := c.Request().Header.Get("conceal-signature")

	//Api Key validation
	if requestApiKey != apiKeyValueConst {
		log.Printf("Invalid API Key")

		return c.JSON(http.StatusUnauthorized, Error{Error: "Invalid API Key"})
	}

	//Timestamp validation
	if !isValidTimestamp(requestTimestamp) {
		log.Printf("Invalid Timestamp")
		return c.JSON(http.StatusBadRequest, Error{Error: "Invalid Timestamp"})
	}

	//Signature validation
	if !isValidSignature(requestTimestamp, requestSignature) {
		log.Printf("Invalid Signature")
		return c.JSON(http.StatusUnauthorized, Error{Error: "Invalid Signature"})
	}

	// Process the webhook payload
	// ..
	logRequest(c)
	// ..

	// Return a success response
	log.Printf("OK")
	return c.JSON(http.StatusOK, Status{Status: "OK"})
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
func logRequest(c echo.Context) error {
	var body map[string]interface{}
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		log.Printf("Invalid Body")
		return c.String(http.StatusUnauthorized, "Invalid Body")
	}
	log.Printf(fmt.Sprintf("req [%s] %s", c.Request().Method, c.Request().URL))
	log.Printf(fmt.Sprintf("headers : %s", c.Request().Header))
	log.Printf(fmt.Sprintf("body: %v", body))
	return nil
}
