from django.urls import path
from . import views

app_name = 'catalog'

urlpatterns = [
    path('', views.index, name='index'),
    path('watch/<int:pk>/', views.detail, name='detail'),
]
