from uuid import uuid4

from django.core.validators import MinValueValidator, MaxValueValidator
from django.db import models
from django.conf import settings

from autoban.common.models import BaseModel
from .validation import VehiclePlateValidator


class Type(BaseModel):
    name = models.CharField(max_length=255, unique=True)

    def __str__(self) -> str:
        return self.name

    class Meta:
        ordering = ['name']


class Brand(BaseModel):
    name = models.CharField(max_length=255)
    type = models.ForeignKey(
        Type,
        on_delete=models.CASCADE,
        related_name='brands'
    )

    class Meta:
        ordering = ['name']
        unique_together = ('type', 'name')

    def __str__(self) -> str:
        return f"{self.name} ({self.type.name})"


class Model(BaseModel):
    type = models.ForeignKey(
        Type,
        on_delete=models.CASCADE,
        related_name='models'
    )
    brand = models.ForeignKey(
        Brand,
        on_delete=models.CASCADE,
        related_name='models'
    )
    name = models.CharField(max_length=255)

    class Meta:
        unique_together = ('brand', 'name')

    def __str__(self) -> str:
        return self.name

    class Meta:
        ordering = ['name']


class Vehicle(BaseModel):
    id = models.UUIDField(primary_key=True, default=uuid4)
    name = models.CharField(max_length=255,
                            null=True,
                            blank=True
                            )
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='vehicles'
    )
    model = models.ForeignKey(
        Model,
        on_delete=models.CASCADE,
        related_name='vehicles'
    )
    color = models.CharField(max_length=255, null=True, blank=True)
    year = models.IntegerField(null=True,
                               blank=True,
                               validators=[
                                   MinValueValidator(1300),
                                   MaxValueValidator(2200)
                               ])
    plate_number = models.CharField(
        validators=[VehiclePlateValidator.validate],
        max_length=255,
        null=True, blank=True
    )
    mileage = models.IntegerField(
        null=True,
        blank=True,
        validators=[
            MinValueValidator(0),
            MaxValueValidator(9999999)
        ])
    insurance_date = models.DateField(null=True, blank=True)

    def __str__(self) -> str:
        if self.name:
            return f"{self.name}"
        return f"Model: {self.model.name}, Plate: {self.plate_number}"

    class Meta:
        ordering = ['name']
