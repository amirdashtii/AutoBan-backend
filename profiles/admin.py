from django.contrib import admin
from .models import Profile

# Define a custom admin class for the User model


class UserAdmin(admin.ModelAdmin):
    # List of fields to display in the admin list view
    list_display = ('user', 'first_name', 'last_name',
                    'birth_date', 'email', 'gender')

    # Enables filter options in the sidebar
    list_filter = ('birth_date', 'gender')

    # Fields to use as search criteria in the admin site
    search_fields = ('first_name', 'last_name', 'email')

    # Fields that are read-only in the detail view
    readonly_fields = ()


# Register the User model along with the custom admin class
admin.site.register(Profile, UserAdmin)
