from django.db.models import Count
from rest_framework import status
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from rest_framework.viewsets import ModelViewSet

from .models import Service, OilChange
from .serializers import (
    ServiceSerializer,
    UpdateServiceSerializer,
    OilChangeSerializer,
    UpdateOilChangeSerializer,
    VehicleServiceSerializer)


class ServiceViewSet(ModelViewSet):
    permission_classes = [IsAuthenticated]

    def get_serializer_class(self):
        if self.request.method == 'PUT':
            return UpdateServiceSerializer
        return ServiceSerializer
    queryset = Service.objects.prefetch_related('oil_change').annotate(
        oil_change_count=Count('oil_change')).all()

    def get_serializer_context(self):
        return {'request': self.request}

    def destroy(self, request, *args, **kwargs):
        if OilChange.objects.filter(service_id=kwargs['pk']).count() > 0:
            return Response({'error': 'Cannot delete service with oil changes'}, status=status.HTTP_400_BAD_REQUEST)
        return super().destroy(request, *args, **kwargs)


class VehicleServiceViewSet(ModelViewSet):

    serializer_class = VehicleServiceSerializer
    permission_classes = [IsAuthenticated]

    def get_queryset(self):
        vehicle_pk = self.kwargs['vehicle_pk']
        return Service.objects.filter(vehicle_id=vehicle_pk)

    def perform_create(self, serializer):
        vehicle_pk = self.kwargs['vehicle_pk']
        serializer.save(vehicle_id=vehicle_pk)


class OilChangeViewSet(ModelViewSet):
    permission_classes = [IsAuthenticated]
    def get_serializer_class(self):
        if self.request.method == 'PUT':
            return UpdateOilChangeSerializer
        return OilChangeSerializer

    def get_queryset(self):
        return OilChange.objects.filter(service_id=self.kwargs['service_pk'])

    def get_serializer_context(self):
        return {'request': self.request}
