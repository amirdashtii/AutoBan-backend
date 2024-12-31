from django.contrib import admin
from .models import Profile


@admin.register(Profile)
class ProfileAdmin(admin.ModelAdmin):
    list_display = ('user', 'first_name', 'last_name',
                    'birth_date', 'gender', 'user__username', 'user__email', 'user__phone_number')
    list_filter = ('birth_date', 'gender', 'user__is_active', 'user__is_admin')
    search_fields = ('first_name', 'last_name',
                     'user__username', 'user__email', 'phone_number')
    readonly_fields = ('user',)

    fieldsets = (
        (None, {
            'fields': ('user', 'first_name', 'last_name', 'birth_date', 'gender', 'user__username', 'user__phone_number', 'user__email')
        }),
        ('Permissions', {
            'fields': ('user__is_active', 'user__is_admin')
        }),
    )
