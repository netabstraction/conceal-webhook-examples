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
	"github.com/labstack/echo/v4/middleware"
)

const signatureKeyConst = "signature-key"
const apiKeyKeyConst = "x-api-key"
const apiKeyValueConst = "sample-key"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + apiKeyKeyConst,
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == apiKeyValueConst, nil
		},
	}))

	e.Use(timestampValidator)
	e.Use(signatureValidator)

	e.Use(middleware.Logger())

	// Routes
	e.POST("/go/echo/api-key-signature-protected", webhookPluginAPI)

	// Start server
	e.Logger.Fatal(e.Start(":4001"))
}

// Exposed protected handler
func webhookPluginAPI(c echo.Context) error {
		logRequest(c, "200 OK")
		return c.String(http.StatusOK, "")
}

// Timestamp validator request timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
func timestampValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp, err := strconv.ParseInt(c.Request().Header.Get("conceal_timestamp"), 10, 64)
		currentTimestamp := time.Now().Unix()
		if err != nil {
			logRequest(c, "Invalid Timestamp")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp")
		}
		log.Println(fmt.Sprintf("Time Diff: %d", requestTimestamp-currentTimestamp))
		if requestTimestamp-currentTimestamp < -60000 || requestTimestamp-currentTimestamp > 120000 {
			logRequest(c, "Invalid Timestamp. Timestamp out of range")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp. Timestamp out of range")
		}
		return next(c)
	}
}

// Signature validator
func signatureValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp := c.Request().Header.Get("conceal_timestamp")
		messageSignature := c.Request().Header.Get("conceal_signature")

		message := fmt.Sprintf("%s|%s://%s%s", requestTimestamp, "http", c.Request().Host, c.Request().URL.Path)
		hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
		hasher.Write([]byte(message))
		sha := fmt.Sprintf("%x", hasher.Sum(nil))

		signatureMatch := sha == messageSignature

		if !signatureMatch {
			logRequest(c, "Signature auth Failed")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusUnauthorized, "Invalid Signature")
		}
		return next(c)
	}
}

func logRequest(c echo.Context, tag string) {
	var body map[string]interface{}
	json.NewDecoder(c.Request().Body).Decode(&body)

	log.Println(fmt.Sprintf("Request Status: %s", tag))
	log.Println("REQUEST")
	log.Println(fmt.Sprintf("Url : %s", c.Request().URL))
	log.Println(fmt.Sprintf("Method : %s", c.Request().Method))
	log.Println(fmt.Sprintf("Header : %s", c.Request().Header))
	log.Println(fmt.Sprintf("Body: %s", body))
}
