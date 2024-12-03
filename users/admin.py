from django.contrib import admin
from .models import User

# Define a custom admin class for the User model


class UserAdmin(admin.ModelAdmin):
    # List of fields to display in the admin list view
    list_display = ('phone_number',  'is_active', 'is_admin',
                    'last_login', 'created_at', 'updated_at')

    # Enables filter options in the sidebar
    list_filter = ('is_active', 'is_admin', 'created_at')

    # Fields to use as search criteria in the admin site
    search_fields = ('phone_number', 'is_active')

    # Fields that are read-only in the detail view
    readonly_fields = ('created_at', 'last_login')


# Register the User model along with the custom admin class
admin.site.register(User, UserAdmin)
