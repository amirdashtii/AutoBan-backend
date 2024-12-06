from django.contrib import admin
from django.db.models.aggregates import Count
from django.db.models.query import QuerySet
from django.http import HttpRequest
from django.urls import reverse
from django.utils.html import format_html
from urllib.parse import urlencode
from . import models


@admin.register(models.Type)
class TypeAdmin(admin.ModelAdmin):
    list_display = ['id', 'name', 'brand_count']
    search_fields = ['name']

    @admin.display(ordering='brand_count')
    def brand_count(self, type):
        url = (
            reverse('admin:vehicles_brand_changelist')
            + '?'
            + urlencode({
                'type__id': str(type.id)
            }))
        return format_html('<a href="{}">{}</a>', url, type.brand_count)

    def get_queryset(self, request):
        return super().get_queryset(request).annotate(brand_count=Count('brands'))


@admin.register(models.Brand)
class BrandAdmin(admin.ModelAdmin):
    autocomplete_fields = ['type']
    list_display = ['id', 'name', 'type', 'type_id', 'model_count']
    list_filter = ['type']
    search_fields = ['name', 'type']

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
        return super().get_queryset(request).annotate(model_count=Count('models'))


@admin.register(models.Model)
class ModelAdmin(admin.ModelAdmin):
    autocomplete_fields = ['brand']
    list_display = ['id', 'name', 'type', 'type_id',
                    'brand', 'brand_id', 'vehicle_count']
    list_filter = ['type', 'brand']
    search_fields = ['name', 'type', 'brand']

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
        return super().get_queryset(request).annotate(vehicle_count=Count('vehicles'))


@admin.register(models.Vehicle)
class VehicleAdmin(admin.ModelAdmin):
    autocomplete_fields = ['user', 'model']
    list_display = ['id', 'name', 'user', 'model', 'plate_number',
                    'color', 'year', 'mileage', 'insurance_date']
    list_filter = ['model']
    search_fields = ['name', 'plate_number',
                     'color', 'year', 'mileage', 'insurance_date']
    ordering = ['id', 'name', 'model', 'plate_number',
                'color', 'year', 'mileage', 'insurance_date']
