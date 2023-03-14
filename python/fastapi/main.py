
from hmac import compare_digest
import time

from fastapi import FastAPI, HTTPException, Request, status

SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"

app = FastAPI()


@app.middleware("api_key_validator")
async def add_api_key_validator(request: Request, call_next):
    print(request.head())
    api_key_header = request.head(API_KEY_CONST)
    print(request.head())
    if (compare_digest(api_key_header, API_KEY_VALUE_CONST)):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid apikey"
        )


@app.middleware("timestamp_validator")
async def add_timestamp_validator(request: Request, call_next):
    start_time = time.time()
    response = await call_next(request)
    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response


@app.post("/python/fastapi/api-key-signature-protected")
def webhook_api():
    print("Hello World")
    return {"Hello": "World"}
