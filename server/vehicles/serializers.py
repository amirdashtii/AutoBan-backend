from rest_framework import serializers
from .models import Type, Brand, Model, Vehicle


class ModelSerializer(serializers.ModelSerializer):
    class Meta:
        model = Model
        fields = ['id', 'name', 'brand']


class BrandSerializer(serializers.ModelSerializer):
    models = ModelSerializer(many=True)

    class Meta:
        model = Brand
        fields = ['id', 'name', 'type', 'model_count', 'models']

    model_count = serializers.IntegerField(read_only=True)


class TypeSerializer(serializers.ModelSerializer):
    brands = BrandSerializer(many=True)

    class Meta:
        model = Type
        fields = ['id', 'name', 'brand_count', 'brands']

    brand_count = serializers.IntegerField(read_only=True)


class VehicleSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)

    class Meta:
        model = Vehicle
        fields = ['id', 'name', 'model', 'color',
                  'plate_number', 'year', 'mileage', 'insurance_date']
