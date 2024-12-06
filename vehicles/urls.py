from django.urls import path, include
from rest_framework_nested import routers
from . import views

# Main Default Router
router = routers.DefaultRouter()
router.register('types', views.TypeViewSet)
router.register('me', views.VehicleViewSet, basename='me')


# Nested Router for Brands under Types
types_router = routers.NestedDefaultRouter(router, 'types', lookup='type')
types_router.register('brands', views.BrandViewSet, basename='type-brands')

# Nested Router for Models under Brands
brands_router = routers.NestedDefaultRouter(
    types_router, 'brands', lookup='brand')
brands_router.register('models', views.ModelViewSet, basename='brand-models')

# Combining URLs from different routers
urlpatterns = [
    path('', include(router.urls)),
    path('', include(types_router.urls)),
    path('', include(brands_router.urls)),
]
