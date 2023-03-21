## Requirement
python (Version 3.11+) 
pip (version 3.11+)

## Build the application
`make build` 

## Run the application
`make run`

## The application exposes following webhook
`http://127.0.0.1:4000/python/django/api-key-signature-protected`

## Webhook Method
POST

## API-KEY-AUTH value
`key` : `x-api-key`

`value` : `sample-key`

## Signature Key value
`signature-key`

## Verification in api
* API Key and value is verified
* conceal_timestamp from request header is verified that is is within range [-60secs, +120secs]
* `conceal_signature` from request header is verified that it is correctly generated. `conceal_signature` is a HMAC signature of the request using SHA256 hashing algoithm. To match the signature, build a string with `conceal_timestamp|webhook_address` That string is then signed with Signature Keyusing SHA256 hashing algoithm.
