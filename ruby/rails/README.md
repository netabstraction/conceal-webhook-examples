## Requirement
ruby 3.2.1

## Build the application
`make build` 

## Run the application
`make run`

## The application exposes following webhook
`http://127.0.0.1:4004/ruby/rails/api-key-signature-protected`

## Webhook Method
POST

## API-KEY-AUTH value
`key` : `x-api-key`

`value` : `sample-key`

## Signature Key value
`signature-key`

## Verification in api
* API Key and value is verified
* conceal-timestamp from request header is verified that is is within range [-60secs, +120secs]
* `cocneal-signature` from request header is verified that it is correctly generated. `conceal-signature` is a HMAC signature of the request using SHA256 hashing algoithm. To match the signature, build a string with `conceal-timestamp|webhook-address` That string is then signed with Signature Keyusing SHA256 hashing algoithm.


## Note
* You may have to use `sudo bundle install` in `make build` errors out with permission issue