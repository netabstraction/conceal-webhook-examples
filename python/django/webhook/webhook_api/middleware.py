import time
import hmac
import hashlib
from django.http import JsonResponse

SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"
WEBHOOK_URL_CONST = "http://127.0.0.1:4000/python/django/api-key-signature-protected"

class APIKeyValidator:
    def __init__(self, get_response):
        self.get_response = get_response
    def __call__(self, request):
        api_key_header = request.headers.get(API_KEY_CONST)
        print(request.headers)
        if api_key_header != API_KEY_VALUE_CONST:
            response = JsonResponse({'error': 'Invalid API Key'})
            response.status_code = 401
            return response
        response = self.get_response(request)
        return response

class SignatureValidator:
    def __init__(self, get_response):
        self.get_response = get_response
    def __call__(self, request):
        request_time = request.headers.get("conceal-timestamp")
        request_signature = request.headers.get("conceal-signature")
        print(request_signature)

        message = '{}|{}'.format(request_time, WEBHOOK_URL_CONST)
        signature = hmac.new(bytes(SIGNATURE_KEY_CONST, 'utf-8'),
                             msg=bytes(message, 'utf-8'), digestmod=hashlib.sha256).hexdigest()
        print(signature)
        if signature != request_signature:
            print("Invalid Signature")
            response = JsonResponse({'error': 'Invalid Signature'})
            response.status_code = 401
            return response
        response = self.get_response(request)
        return response

class TimestampValidator:
    def __init__(self, get_response):
        self.get_response = get_response
    def __call__(self, request):
        current_time = int(time.time())
        request_time = int(request.headers.get("conceal-timestamp"))

        if (request_time - current_time < -60000 or request_time - current_time > 12000):
            print("Invalid Timestamp. Timestamp not in range")
            response = JsonResponse({'error': 'Invalid Timestamp. Timestamp not in range'})
            response.status_code = 400
            return response
        response = self.get_response(request)
        return response

class LoggerMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response
    def __call__(self, request):
        print("REQUEST")
        print("Url: ", request.build_absolute_uri())
        print("Method: ", request.method)
        print("Header: ", request.headers)
        print("Body: ", request.body)
        return self.get_response(request)
