from django.urls import path, include

urlpatterns = [
    path('auth/', include(('djoser.urls', 'auth'))),
    path('auth/', include(('djoser.urls.jwt', 'jwt_auth'))),
    # path('users/', include(('users.urls', 'users'))),
    path('profiles/', include(('profiles.urls', 'profiles'))),
    path('vehicles/', include(('vehicles.urls', 'vehicles'))),
    path('services/', include(('services.urls', 'services'))),
]
