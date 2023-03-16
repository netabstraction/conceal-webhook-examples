import time
import hmac
import hashlib
import json

from hmac import compare_digest
from fastapi import FastAPI, HTTPException, Request
from fastapi.responses import JSONResponse


SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"
WEBHOOK_URL_CONST = "http://127.0.0.1:4004/python/fastapi/api-key-signature-protected"

app = FastAPI()


@app.middleware("api_key_validator")
async def add_api_key_validator(request: Request, call_next):
    api_key_header = request.headers.get(API_KEY_CONST)
    if not (compare_digest(api_key_header, API_KEY_VALUE_CONST)):
        print("Invalid Key")
        return JSONResponse(content="Invalid Key", status_code=401)

    return await call_next(request)


@app.middleware("signature_validator")
async def add_timestamp_validator(request: Request, call_next):

    request_time = int(request.headers.get("conceal_timestamp"))
    request_signature = request.headers.get("conceal_signature")
    message = '{}|{}'.format(request_time, WEBHOOK_URL_CONST)
    signature = hmac.new(bytes(SIGNATURE_KEY_CONST, 'utf-8'),
                         msg=bytes(message, 'utf-8'), digestmod=hashlib.sha256).hexdigest()
    if signature != request_signature:
        print("Invalid Signature")
        return JSONResponse(content="Invalid Signature", status_code=401)

    return await call_next(request)


@app.middleware("timestamp_validator")
async def add_timestamp_validator(request: Request, call_next):
    current_time = int(time.time())
    request_time = int(request.headers.get("conceal_timestamp"))
    if (request_time - current_time < -60000 or request_time - current_time > 12000):
        print("Invalid Timestamp. Timestamp not in range")
        return JSONResponse(content="Invalid Timestamp. Timestamp not in range", status_code=400)

    return await call_next(request)


@app.middleware("logger")
async def add_api_key_validator(request: Request, call_next):
    print("REQUEST")
    print("Url: ", request.url)
    print("Method: ", request.method)
    print("Header: ", request.headers)
    print("Body: ", await request.json())
    return await call_next(request)


@app.post("/python/fastapi/api-key-signature-protected")
async def webhook_api():
    print("200 OK")
    return {"ok": "true"}
