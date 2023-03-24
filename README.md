# conceal-webhook-examples
Webhook examples verify following information for validation. 

## Verification in api
* API Key and value is verified
* conceal-timestamp from request header is verified that is is within range [-60secs, +120secs]
* `cocneal-signature` from request header is verified that it is correctly generated. `conceal-signature` is a HMAC signature of the request using SHA256 hashing algoithm. To match the signature, build a string with `conceal-timestamp|webhook-address` That string is then signed with Signature Keyusing SHA256 hashing algoithm.

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
* python
  * django
  * fastapi


## Testing your webhook using postman

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

Check this exported [Postman Example](../test-util/Webhook_Example.postman_collection.json)