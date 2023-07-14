import time
import hmac
import hashlib

from hmac import compare_digest
from flask import Flask, Response, request, jsonify

SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"
WEBHOOK_URL_CONST = "http://127.0.0.1:8080/webhook"

app = Flask(__name__)

@app.route('/webhook', methods=['POST'])
async def handle_webhook():
    request_timestamp = request.headers.get("conceal-timestamp")
    request_signature = request.headers.get("conceal-signature")
    request_api_key = request.headers.get(API_KEY_CONST)

    # API Key Validation
    if not (compare_digest(request_api_key, API_KEY_VALUE_CONST)):
        print("Invalid Key")
        return jsonify({"error": "Invalid API Key"}), 401

    # Timestamp Validation
    if not(is_valid_timestamp(request_timestamp)):
        print("Invalid Timestamp")
        return jsonify({"error": "Invalid Timestamp"}), 400

    # Signature Validation
    if not(is_valid_signature(request_timestamp, request_signature)):
        print("Invalid Signature")
        return jsonify({"error": "Invalid Signature"}), 401

    # Process the webhook payload
    # ..
    await log_request(request)

    print("OK")
    return jsonify({"status": "OK"})

def is_valid_timestamp(request_timestamp):
    if(request_timestamp is None):
        return False
    try:
        request_timestamp_int = int(request_timestamp)
    except:
        return False

    current_timestamp = int(time.time())
    return (request_timestamp_int - current_timestamp > -60000 and request_timestamp_int - current_timestamp < 12000)

def is_valid_signature(request_timestamp, request_signature):
    message = '{}|{}'.format(request_timestamp, WEBHOOK_URL_CONST)
    signature = hmac.new(bytes(SIGNATURE_KEY_CONST, 'utf-8'),
                         msg=bytes(message, 'utf-8'), digestmod=hashlib.sha256).hexdigest()
    return signature == request_signature

async def log_request(request):
    print("req [{method}] {url}".format(method = request.method, url = request.url))
    print("headers: ", request.headers)
    print("body: ", request.json)

app.run(host='0.0.0.0', port=8080)
