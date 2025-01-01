from uuid import uuid4

from django.conf import settings
from django.core.validators import MinValueValidator, MaxValueValidator
from django.db import models
from django.utils import timezone

from autoban.common.models import BaseModel
from vehicles.models import Vehicle
from .validation import future_date_validator


class Service(BaseModel):
    id = models.UUIDField(primary_key=True, default=uuid4)
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='services'
    )
    vehicle = models.ForeignKey(
        Vehicle,
        on_delete=models.CASCADE,
        related_name='services')
    service_date = models.DateField(validators=[future_date_validator])
    mileage = models.IntegerField(
        validators=[
            MinValueValidator(0),
            MaxValueValidator(9999999)
        ])
    description = models.TextField(blank=True, null=True)

    def __str__(self):
        return str(self.id)


class OilChange(BaseModel):
    service = models.OneToOneField(
        Service,
        on_delete=models.CASCADE,
        related_name='oil_change')
    oil_type = models.CharField(max_length=255)
    oil_lifetime_distance = models.PositiveIntegerField(blank=True, null=True)
    next_change_mileage = models.PositiveIntegerField(blank=True, null=True)
    next_service_date = models.DateField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)

    def __str__(self):
        return f"Oil Change: {self.oil_type}"
