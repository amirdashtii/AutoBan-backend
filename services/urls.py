from django.urls import path, include
from rest_framework_nested import routers

from . import views

# Main Default Router
router = routers.DefaultRouter()
router.register('', views.ServiceViewSet, basename='services')
router.register('vehicles/(?P<vehicle_pk>[^/.]+)',
                views.VehicleServiceViewSet, basename='vehicle_service')


services_router = routers.NestedDefaultRouter(router, '', lookup='service')
services_router.register(
    'oil_change', views.OilChangeViewSet, basename='service-oil_change')


# Combining URLs from different routers
urlpatterns = [
    path('', include(router.urls)),
    path('', include(services_router.urls)),
]
