from django.contrib import admin
from django.db.models.aggregates import Count
from django.db.models.query import QuerySet
from django.http import HttpRequest
from django.urls import reverse
from django.utils.html import format_html
from urllib.parse import urlencode
from . import models


@admin.register(models.VehicleType)
class VehicleTypeAdmin(admin.ModelAdmin):
    list_display = ['name', 'brand_count']
    ordering = ['name']
    search_fields = ['name']

    @admin.display(ordering='brand_count')
    def brand_count(self, vehicle_type):
        url = (
            reverse('admin:vehicles_brand_changelist')
            + '?'
            + urlencode({
                'vehicle_type__id': str(vehicle_type.id)
            }))
        return format_html('<a href="{}">{}</a>', url, vehicle_type.brand_count)

    def get_queryset(self, request):
        return super().get_queryset(request).annotate(brand_count=Count('brand'))


@admin.register(models.Brand)
class BrandAdmin(admin.ModelAdmin):
    autocomplete_fields = ['vehicle_type']
    list_display = ['name', 'vehicle_type', 'model_count']
    list_filter = ['vehicle_type']
    ordering = ['name', 'vehicle_type']
    search_fields = ['name']

    @admin.display(ordering='model_count')
    def model_count(self, brand):
        url = (
            reverse('admin:vehicles_model_changelist')
            + '?'
            + urlencode({
                'brand__id': str(brand.id)
            }))
        return format_html('<a href="{}">{}</a>', url, brand.model_count)

    def get_queryset(self, request):
        return super().get_queryset(request).annotate(model_count=Count('model'))


@admin.register(models.Model)
class ModelAdmin(admin.ModelAdmin):
    autocomplete_fields = ['brand']
    list_display = ['name', 'brand', 'vehicle_count']
    list_filter = ['brand']
    ordering = ['name', 'brand']
    search_fields = ['name']

    @admin.display(ordering='vehicle_count')
    def vehicle_count(self, model):
        url = (
            reverse('admin:vehicles_vehicle_changelist')
            + '?'
            + urlencode({
                'model__id': str(model.id)
            }))
        return format_html('<a href="{}">{}</a>', url, model.vehicle_count)

    def get_queryset(self, request: HttpRequest) -> QuerySet:
        return super().get_queryset(request).annotate(vehicle_count=Count('vehicle'))


@admin.register(models.Vehicle)
class VehicleAdmin(admin.ModelAdmin):
    autocomplete_fields = ['user', 'model']
    list_display = ['name', 'user', 'model', 'plate_number',
                    'color', 'year', 'mileage', 'insurance_date']
    list_filter = ['model']
    search_fields = ['name', 'plate_number',
                     'color', 'year', 'mileage', 'insurance_date']
    ordering = ['name', 'model', 'plate_number',
                'color', 'year', 'mileage', 'insurance_date']
