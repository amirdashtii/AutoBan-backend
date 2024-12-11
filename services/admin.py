from django.contrib import admin
from . import models


@admin.register(models.Service)
class ServiceAdmin(admin.ModelAdmin):
    autocomplete_fields = ['user', 'vehicle']
    list_display = ['id', 'user', 'vehicle',
                    'service_date', 'mileage', 'description']
    search_fields = ['user', 'vehicle', 'service_date', 'mileage']


@admin.register(models.OilChange)
class OilChangeAdmin(admin.ModelAdmin):
    autocomplete_fields = ['service']
    list_display = ['service', 'oil_type', 'oil_lifetime_distance',
                    'next_change_mileage', 'next_service_date', 'description']
    list_filter = ['oil_type']
    search_fields = ['oil_type']
