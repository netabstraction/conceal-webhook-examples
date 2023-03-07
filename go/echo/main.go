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

	e.Use(TimestampValidator)
	e.Use(SignatureValidator)

	e.Use(middleware.Logger())

	// Routes
	e.GET("/go/echo/api-key-signature-protected", protectedHandler)

	// Start server
	e.Logger.Fatal(e.Start(":4001"))
}

// Exposed protected handler
func protectedHandler(c echo.Context) error {
		log.Println("This is the protected handler")
		return c.String(http.StatusOK, "Hello, World!")
}

// Timestamp validator
func TimestampValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp, err := strconv.ParseInt(c.Request().Header.Get("conceal_timestamp"), 10, 64)
		currentTimestamp := time.Now().Unix()
		if err != nil {
			LogRequest(c, "Invalid Timestamp")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp")
		}
		if requestTimestamp-currentTimestamp > 60000 || currentTimestamp-requestTimestamp > 120000 {
			LogRequest(c, "Invalid Timestamp. Timestamp out of range")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp. Timestamp out of range")
		}
		return next(c)
	}
}

// Signature validator
func SignatureValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp := c.Request().Header.Get("conceal_timestamp")
		messageSignature := c.Request().Header.Get("conceal_signature")

		message := fmt.Sprintf("%s|%s://%s%s", requestTimestamp, "http", c.Request().Host, c.Request().URL)
		hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
		hasher.Write([]byte(message))
		sha := fmt.Sprintf("%x", hasher.Sum(nil))

		signatureMatch := sha == messageSignature

		if !signatureMatch {
			LogRequest(c, "Signature auth Failed")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusUnauthorized, "Invalid Signature")
		}
		return next(c)
	}
}

func LogRequest(c echo.Context, tag string) {
	log.Println("Got a new request " + tag)

	log.Println(c.Request().URL)
	log.Println(c.Request().Method)
	log.Println(c.Request().Header)
	log.Println(c.Request().URL)

	var body map[string]interface{}
	json.NewDecoder(c.Request().Body).Decode(&body)

	log.Println(body)
}