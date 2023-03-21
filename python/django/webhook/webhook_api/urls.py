from django.urls import path, include
from .views import WebhookApiView

urlpatterns = [
    path('python/django/api-key-signature-protected', WebhookApiView.as_view()),
]
