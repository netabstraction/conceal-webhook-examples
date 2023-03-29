from django.urls import path, include
from .views import WebhookApiView

urlpatterns = [
    path('webhook', WebhookApiView.as_view()),
]
