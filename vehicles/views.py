from django.db.models import Count

from rest_framework import status
from rest_framework.exceptions import NotAuthenticated
from rest_framework.permissions import IsAdminUser, AllowAny, IsAuthenticated
from rest_framework.response import Response
from rest_framework.viewsets import ModelViewSet


from .models import Type, Brand, Model, Vehicle
from .serializers import TypeSerializer, BrandSerializer, ModelSerializer, VehicleSerializer


class TypeViewSet(ModelViewSet):
    queryset = Type.objects.annotate(brand_count=Count('brands')).all()
    serializer_class = TypeSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAdminUser()]


class BrandViewSet(ModelViewSet):
    serializer_class = BrandSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAdminUser()]

    def get_queryset(self):
        return Brand.objects.filter(type_id=self.kwargs['type_pk'])

    def get_serializer_context(self):
        return {'type_id': self.kwargs['type_pk']}

    def destroy(self, request, *args, **kwargs):
        if Model.objects.filter(brand_id=kwargs['pk']).count() > 0:
            return Response({'error': 'Cannot delete brand with models'}, status=status.HTTP_400_BAD_REQUEST)
        return super().destroy(request, *args, **kwargs)


class ModelViewSet(ModelViewSet):
    serializer_class = ModelSerializer

    def get_permissions(self):
        if self.request.method == 'GET':
            return [AllowAny()]
        return [IsAdminUser()]

    def get_serializer_context(self):
        return {'request': self.request}

    def get_queryset(self):
        return Model.objects.filter(
            brand_id=self.kwargs['brand_pk'],
            type_id=self.kwargs['type_pk']
        )

    def destroy(self, request, *args, **kwargs):
        if Vehicle.objects.filter(model_id=kwargs['pk']).count() > 0:
            return Response({'error': 'Cannot delete model with vehicles'}, status=status.HTTP_400_BAD_REQUEST)
        return super().destroy(request, *args, **kwargs)


class VehicleViewSet(ModelViewSet):
    serializer_class = VehicleSerializer
    permission_classes = [IsAuthenticated]

    def get_queryset(self):
        user = self.request.user
        if user.is_authenticated:
            return Vehicle.objects.filter(user=user)
        else:
            raise NotAuthenticated(
                "Authentication credentials were not provided.")
