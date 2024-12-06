from rest_framework import serializers
from .models import Type, Brand, Model, Vehicle


class TypeSerializer(serializers.ModelSerializer):
    class Meta:
        model = Type
        fields = ['id', 'name', 'brand_count']

    brand_count = serializers.IntegerField(read_only=True)


class BrandSerializer(serializers.ModelSerializer):
    class Meta:
        model = Brand
        fields = ['id', 'name', 'type', 'model_count']

    model_count = serializers.IntegerField(read_only=True)


class ModelSerializer(serializers.ModelSerializer):
    class Meta:
        model = Model
        fields = ['id', 'name', 'brand']


class VehicleSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)

    class Meta:
        model = Vehicle
        fields = ['id', 'name', 'model', 'color',
                  'plate_number', 'year', 'mileage', 'insurance_date']
