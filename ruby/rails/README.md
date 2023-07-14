Requirements
--

* ruby 3.2.1
* rails 7.0.6

Build
--

```bash
make build
```

Run
--

```bash
make run
```

Application exposes webhook endpoint at address
`http://127.0.0.1:8080/webhook`

Test
--

The example webhook can be tested using the included Postman collection `Webhook Example.postman_collection.json`.

Verification
--

To validate the webhook is functioning properly ensure:

1. The request returns a valid response

```bash
{
    "status": "OK"
}
```

2. The API key is being validated. If the `x-api-key` header does not match the expected `sample-key` value the following `400 Bad Request` error should be returned.

```bash
{
    "error": "Invalid API Key"
}
```

3. The signature is being validated. The `conceal-signature` header is an `HMAC` signature calculated using the string `<conceal-timestamp>|<webhook_address>` which is then signed using the value `signature-key` as the Signature Key for the `SHA256` hashing algorithm.

   If the `conceal-signature` does not match the value calculated by the webhook the following `401 Unauthorized` error should be returned.

```bash
{
    "error": "Invalid Signature"
}
```

4. The timestamp is being validated. If the `conceal-timestamp` header is not within the range `[-60 secs, +120 secs]` then the following `400 Bad Request` error should be returned.

```bash
{
    "error": "Invalid Timestamp"
}
```
