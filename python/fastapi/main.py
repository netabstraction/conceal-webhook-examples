from pickle import FALSE
import time
import hmac
import hashlib

from hmac import compare_digest
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse


SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"
WEBHOOK_URL_CONST = "http://127.0.0.1:8080/webhook"

app = FastAPI()

@app.post("/webhook")
async def webhook_api(request: Request):

    request_timestamp = request.headers.get("conceal-timestamp")
    request_signature = request.headers.get("conceal-signature")
    request_api_key = request.headers.get(API_KEY_CONST)

    # Api Key validation
    if not (compare_digest(request_api_key, API_KEY_VALUE_CONST)):
        print("Invalid Key")
        return JSONResponse(content="Invalid Key", status_code=401)

    # Timestamp validation
    if not (is_valid_timestamp(request_timestamp)):
        print("Invalid Timestamp")
        return JSONResponse(content="Invalid Timestamp", status_code=400)

    # Signature validation
    if not (is_valid_signature(request_timestamp, request_signature)):
        print("Invalid Signature")
        return JSONResponse(content="Invalid Signature", status_code=401)

    # Validate json body
    try:
        body = await request.json()
    except:
        print("Invalid Body")
        return JSONResponse(content="Invalid Body", status_code=400)
   
    # Process the webhook payload
    # ..
    await request_log(request)
    # ..

    # Return a success response
    print("OK")
    return JSONResponse(content="", status_code=200)

# Validate timestamp timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
def is_valid_timestamp(request_timestamp: any):
    if(request_timestamp is None):
        return False
    
    try:
        request_timestamp_int = int(request_timestamp)
    except:
        return False

    current_timestamp = int(time.time())
    return (request_timestamp_int - current_timestamp > -60000 and request_timestamp_int - current_timestamp < 12000)

# Validate signature
def is_valid_signature(request_timestamp: any, request_signature: any):
    message = '{}|{}'.format(request_timestamp, WEBHOOK_URL_CONST)
    signature = hmac.new(bytes(SIGNATURE_KEY_CONST, 'utf-8'),
                         msg=bytes(message, 'utf-8'), digestmod=hashlib.sha256).hexdigest()
    return signature == request_signature

# Log request
async def request_log(request: Request):
    print("req [{method}] {url}".format(method = request.method, url = request.url))
    print("headers: ", request.headers)
    print("body: ", await request.json())
