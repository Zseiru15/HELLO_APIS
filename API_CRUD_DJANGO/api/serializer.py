from rest_framework import serializers
from .models import programmer
from .models import student

class ProgrammerSerializer(serializers.ModelSerializer):
    class Meta:
        model=programmer
        #fields= ('full_name', 'nickname', 'age') #ejemplo largo o selectivo
        fields= '__all__' #ejemplo corto o seleccion completa
        
class StudentSerializer(serializers.ModelSerializer):
    class Meta:
        model=student
        #fields= ('full_name', 'nickname', 'age') #ejemplo largo o selectivo
        fields= '__all__' #ejemplo corto o seleccion completa