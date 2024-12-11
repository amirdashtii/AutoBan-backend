from rest_framework import serializers
from .models import Service, OilChange


class OilChangeSerializer(serializers.ModelSerializer):
    class Meta:
        model = OilChange
        fields = ['id', 'service', 'oil_type', 'oil_lifetime_distance', 'next_change_mileage',
                  'next_service_date', 'description']


class UpdateOilChangeSerializer(serializers.ModelSerializer):
    class Meta:
        model = OilChange
        fields = ['oil_type', 'oil_lifetime_distance', 'next_change_mileage',
                  'next_service_date', 'description']


class ServiceSerializer(serializers.ModelSerializer):
    id = serializers.UUIDField(read_only=True)
    oil_change = OilChangeSerializer(read_only=True)

    class Meta:
        model = Service
        fields = ['id', 'user', 'vehicle', 'service_date',
                  'mileage', 'description', 'oil_change']


class UpdateServiceSerializer(serializers.ModelSerializer):
    class Meta:
        model = Service
        fields = ['service_date', 'mileage', 'description']


