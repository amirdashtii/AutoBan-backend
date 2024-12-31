from django.contrib import admin
from . import models


@admin.register(models.Service)
class ServiceAdmin(admin.ModelAdmin):
    autocomplete_fields = ['vehicle']
    list_display = ['id', 'user', 'vehicle',
                    'service_date', 'mileage', 'description']
    search_fields = ['user__username',
                     'vehicle__name', 'service_date', 'mileage']
    readonly_fields = ['id']
    list_filter = ['service_date', 'vehicle__name']

    fieldsets = (
        (None, {
            'fields': ('user', 'vehicle', 'service_date', 'mileage', 'description')
        }),
        ('Advanced options', {
            'classes': ('collapse',),
            'fields': ('id',)
        }),
    )


@admin.register(models.OilChange)
class OilChangeAdmin(admin.ModelAdmin):
    autocomplete_fields = ['service']
    list_display = ['service', 'oil_type', 'oil_lifetime_distance',
                    'next_change_mileage', 'next_service_date', 'description']
    list_filter = ['oil_type', 'next_service_date']
    search_fields = ['oil_type', 'service__user__phone_number',
                     'service__user__username', 'service__vehicle__name']
    readonly_fields = ['id']

    fieldsets = (
        (None, {
            'fields': ('service', 'oil_type', 'oil_lifetime_distance', 'next_change_mileage', 'next_service_date', 'description')
        }),
        ('Advanced options', {
            'classes': ('collapse',),
            'fields': ('id',)
        }),
    )
