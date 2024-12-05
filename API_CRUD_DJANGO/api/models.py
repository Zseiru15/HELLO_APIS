from django.db import models

# Create your models here.

class programmer(models.Model):
    fullname = models.CharField(max_length=50)
    nickname = models.CharField(max_length=30)
    age = models.IntegerField(default=True)
    phone = models.CharField(max_length=10, null=True, default=None)
    is_active = models.BooleanField(default=True)
    
class student(models.Model):
    name = models.CharField(max_length=20)
    lastname = models.CharField(max_length=20)
    gender = models.CharField(max_length=7)
    token_number = models.CharField(max_length=7, null=True, default=None)
    training = models.BooleanField(default=True)
    entry_date = models.IntegerField(default=True)
    is_active = models.BooleanField(default=True)