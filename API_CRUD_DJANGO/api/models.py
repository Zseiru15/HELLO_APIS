from django.db import models

# Create your models here.

class programmer(models.Model):
    fullname = models.CharField(max_length=50)
    nickname = models.CharField(max_length=30)
    age = models.IntegerField(default=True)
    phone = models.CharField(max_length=10, null=True, default=None)
    is_active = models.BooleanField(default=True)