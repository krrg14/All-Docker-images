from django.contrib import admin
from .models import Watch

@admin.register(Watch)
class WatchAdmin(admin.ModelAdmin):
    list_display = ('name', 'price')
