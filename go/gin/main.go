package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type application struct {
}

const (
	signatureKeyConst = "signature-key"
	apiKeyValueConst  = "sample-key"
	apiKeyKeyConst    = "x-api-key"
	webhookAddress    = "http://127.0.0.1:8080/webhook"
)

func main() {
	router := gin.New()

	// routes
	router.POST("/webhook", handleWebhook)

	log.Printf("starting server on %s", ":8080")
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("127.0.0.1:8080")
}

func handleWebhook(c *gin.Context) {
	requestApiKey := c.Request.Header.Get(apiKeyKeyConst)
	requestTimestamp := c.Request.Header.Get("conceal-timestamp")
	requestSignature := c.Request.Header.Get("conceal-signature")

	// API Key Validation
	if requestApiKey != apiKeyValueConst {
		log.Printf("Invalid API Key")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid API Key"})
		c.Abort()
		return
	}

	// Timestamp Validation
	if !isValidTimestamp(requestTimestamp) {
		log.Printf("Invalid Timestamp")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "Invalid Timestamp"})
		c.Abort()
		return
	}

	// Signature Validation
	if !isValidSignature(requestTimestamp, requestSignature) {
		log.Printf("Invalid Signature")
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"err": "Invalid Signature"})
		c.Abort()
		return
	}

	logRequest(c)

	log.Printf("OK")
	c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
	c.Next()
}

// Validate timestamp. Timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
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

// Validate Signature
func isValidSignature(requestTimestamp string, requestSignature string) bool {
	message := fmt.Sprintf("%s|%s", requestTimestamp, webhookAddress)
	hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
	hasher.Write([]byte(message))
	sha := fmt.Sprintf("%x", hasher.Sum(nil))

	return sha == requestSignature
}

func logRequest(c *gin.Context) {
	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}
