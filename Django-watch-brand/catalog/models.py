# This is a Django model class for a Watch with fields for name, description, price, and image URL.
from django.db import models

class Watch(models.Model):
    name = models.CharField(max_length=200)
    description = models.TextField(blank=True)
    price = models.DecimalField(max_digits=10, decimal_places=2)
    image_url = models.URLField(blank=True)

    def __str__(self):
        return self.name
