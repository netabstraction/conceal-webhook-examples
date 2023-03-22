from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions

# Create your views here.
class WebhookApiView(APIView):
    #permission_classes = [permissions.IsAuthenticated]

    def post(self, request, *args, **kwargs):
        return Response({"msg": "ok"}, status=status.HTTP_200_OK, template_name=None, headers=None, content_type=None)
