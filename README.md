# conceal-webhook-examples
Webhook examples verify following information for validation. 

## List of examples
* c#
  * netcore
* go
  * echo
  * gin
  * nethttp
* java
  * spring
* nodejs
  * express
  * express-ts
  * fastify
* python
  * django
  * fastapi
  * flask
* ruby
  * rails

## Testing your webhook using Postman

Add following script to prerequest script
```
const currentTimestamp = Math.floor(Date.now() / 1000);
const webhookUrl = postman.getGlobalVariable("webhookUrl");
const signatureKey = postman.getGlobalVariable("signature-key")
const message = `${currentTimestamp}|${webhookUrl}`;
const signature = CryptoJS.HmacSHA256(message, signatureKey);

postman.setGlobalVariable("conceal-signature", signature );
postman.setGlobalVariable("conceal-timestamp", currentTimestamp);
```
* Set `signature-key` and `webhookUrl` in environment variable

* Sample Body

```
{
    "event": "Scanned URL",
    "host": "Conceal API",
    "sourcetype": "Conceal API Post Process",
    "company_id": "91140740-fe8a-4350-ad74-d79bb2828318",
    "company_name": "Sample Company",
    "user_email": "user1@sample_company.io",
    "user_id": "e2bfb697-a3c6-429a-89fc-ffef546d348b",
    "url": "https://www.google.com",
    "count": 2,
    "decision": {
      "enforcetls": "allow",
      "noip": "allow"
    },
    "final_decision": "allow",
    "timestamp": 1678113806
}
```

Alternatively, use this exported [Postman Example](../test-util/Webhook_Example.postman_collection.json)

## Verification in api
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
