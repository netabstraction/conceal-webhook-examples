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

const signatureKeyConst = "signature-key"
const apiKeyValueConst = "sample-key"

func main() {
	router := gin.New()

	// signature middleware
	router.Use(signatureAuth)
	// api key auth
	router.Use(apiKeyAuth)

	// routes
	router.POST("/go/gin/api-key-signature-protected", protectedHandler)

	log.Printf("starting server on %s", ":4000")
	router.Run("localhost:4000")
}

// Exposed protected handler
func protectedHandler(c *gin.Context) {
	logRequest(c, "200 OK")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "works"})
	c.Next()
}

func apiKeyAuth(c *gin.Context) {
	apiKeyMatch := c.Request.Header.Get("x-api-key") == apiKeyValueConst

	if !apiKeyMatch {
		logRequest(c, "Api key auth failed")
		c.Writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid API Key"})
		c.Abort()
		return
	}

	requestTimestamp, err := strconv.ParseInt(c.Request.Header.Get("conceal_timestamp"), 10, 64)
	if err != nil {
		logRequest(c, "Invalid timestamp")
		c.Writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp"})
		c.Abort()
		return
	}

	currentTimestamp := time.Now().Unix()
	log.Println(fmt.Sprintf("Time Diff: %d", requestTimestamp-currentTimestamp))

	if requestTimestamp-currentTimestamp < -60000 || requestTimestamp-currentTimestamp > 120000 {
		logRequest(c, "Invalid timestamp. Timestamp out of range")
		c.Writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp. Timestamp out of range"})
		c.Abort()
		return
	}
}

func signatureAuth(c *gin.Context) {
	timestamp := c.Request.Header.Get("conceal_timestamp")
	messageSignature := c.Request.Header.Get("conceal_signature")

	message := fmt.Sprintf("%s|%s://%s%s", timestamp, "http", c.Request.Host, c.Request.URL.Path)
	hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
	hasher.Write([]byte(message))
	sha := fmt.Sprintf("%x", hasher.Sum(nil))

	signatureMatch := sha == messageSignature

	if !signatureMatch {
		fmt.Printf("message: %s\r\n", message)
		fmt.Printf("signature: %s\r\n", messageSignature)
		fmt.Printf("sha: %s\r\n", sha)
		logRequest(c, "Signature auth failed")
		c.Writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
		c.Abort()
		return
	}
}

func logRequest(c *gin.Context, tag string) {
	log.Println("Got a new request " + tag)

	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}
