## Build the application
`make build`

## Run the application
`make run `

## The application exposes following webhook
`http://127.0.0.1:8080/webhook

## API-KEY-AUTH value
`key` : `x-api-key`

`value` : `sample-key`

## Signature Key value
`signature-key`

## Verification in api
* API Key and value is verified
* conceal-timestamp from request header is verified that is is within range [-60secs, +120secs]
* `conceal-signature` from request header is verified that it is correctly generated. `conceal-signature` is a HMAC signature of the request using SHA256 hashing algoithm. To match the signature, build a string with `conceal-timestamp|webhook_address` That string is then signed with Signature Key using SHA256 hashing algorithm.
