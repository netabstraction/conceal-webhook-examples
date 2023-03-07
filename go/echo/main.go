package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
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
func protectedHandler(c echo.Context) echo.HandlerFunc {
	// app.logRequest(c, "Protected Handler")
	return func(c echo.Context) error {
		log.Println("This is the protected handler")
		return c.String(http.StatusOK, "Hello, World!")
	}
}

// Timestamp validator
func TimestampValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp, err := strconv.ParseInt(c.Request().Header.Get("conceal_timestamp"), 10, 64)
		currentTimestamp := time.Now().Unix()
		if err != nil {
			//app.logRequest(c, "Invalid Timestamp")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp")
		}
		if requestTimestamp-currentTimestamp > 60000 || currentTimestamp-requestTimestamp > 120000 {
			//app.logRequest(c, "Invalid Timestamp. Timestamp out of range")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusBadRequest, "Invalid Timestamp. Timestamp out of range")
		}
		return next(c)
	}
}

// Signature validator
func SignatureValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTimestamp, err := c.Request().Header.Get("conceal_timestamp")
		messageSignature := c.Header.Get("conceal_signature")

		message := fmt.Sprintf("%s|%s://%s%s", requestTimestamp, "http", c.Request().Host, c.Request().Path)
		hasher := hmac.New(sha256.New, []byte(signatureKeyConst))
		hasher.Write([]byte(message))
		sha := fmt.Sprintf("%x", hasher.Sum(nil))

		signatureMatch := sha == messageSignature

		if !signatureMatch {
			//app.logRequest(c, "Signature auth Failed")
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			return c.String(http.StatusUnauthorized, "Invalid Signature")
		}
		return next(c)
	}
}

// func logRequest(c echo.Context, tag string) {
// 	log.Println("Got a new request " + tag)

// 	log.Println(c.Request().URL)
// 	// log.Println(r.Method)
// 	// log.Println(r.Header)
// 	// log.Println(r.URL.RawQuery)

// 	var body map[string]interface{}
// 	// json.NewDecoder(r.Body).Decode(&body)

// 	log.Println(body)
// }
