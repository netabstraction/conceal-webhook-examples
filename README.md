# conceal-webhook-examples


# Testing your webhook 

## Postman

Check this [exported postman](./test-util/Webhook Example.postman_collection.json)

* Add following script to prerequest script
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