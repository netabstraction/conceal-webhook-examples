import time
import hmac
import hashlib

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions

from django.http import JsonResponse

SIGNATURE_KEY_CONST = "signature-key"
API_KEY_CONST = "x-api-key"
API_KEY_VALUE_CONST = "sample-key"
WEBHOOK_URL_CONST = "http://127.0.0.1:8080/webhook"

# Create your views here.
class WebhookApiView(APIView):
    #permission_classes = [permissions.IsAuthenticated]

    def post(self, request, *args, **kwargs):
        requestApiKey = request.headers.get(API_KEY_CONST)
        requestTimestamp = int(request.headers.get("conceal-timestamp"))
        requestSignature = request.headers.get("conceal-signature")

        if requestApiKey != API_KEY_VALUE_CONST:
            response = JsonResponse({'error': 'Invalid API Key'})
            response.status_code = 401
            return response

        if not isValidTimestamp(requestTimestamp):
            print("Invalid Timestamp. Timestamp not in range")
            response = JsonResponse({'error': 'Invalid Signature'})
            response.status_code = 400
            return response

        if not isValidSignature(requestTimestamp, requestSignature):
            print("Invalid Signature")
            response = JsonResponse({'error': 'Invalid Signature'})
            response.status_code = 401
            return response

        logRequest(request)

        return Response({"msg": "ok"}, status=status.HTTP_200_OK, template_name=None, headers=None, content_type=None)

def isValidTimestamp(requestTimestamp):
    currentTimestamp = int(time.time())
    
    if (requestTimestamp - currentTimestamp < -60000) or (requestTimestamp - currentTimestamp > 120000):
        return False
    return True

def isValidSignature(requestTimestamp, requestSignature):
    message = '{}|{}'.format(requestTimestamp, WEBHOOK_URL_CONST)
    signature = hmac.new(bytes(SIGNATURE_KEY_CONST, 'utf-8'),
                         msg=bytes(message, 'utf-8'), digestmod=hashlib.sha256).hexdigest()

    if signature != requestSignature:
        return False
    return True

def logRequest(request):
    print("Url: ", request.build_absolute_uri())
    print("Method: ", request.method)
    print("Header: ", request.headers)
    print("Body: ", request.body)
