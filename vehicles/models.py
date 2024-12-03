from django.core.validators import MinValueValidator, MaxValueValidator
from django.db import models
from django.conf import settings

from .validation import VehiclePlateValidator


class VehicleType(models.Model):
    name = models.CharField(max_length=255, unique=True)

    def __str__(self) -> str:
        return self.name

    class Meta:
        ordering = ['name']


class Brand(models.Model):
    vehicle_type = models.ForeignKey(
        VehicleType,
        on_delete=models.CASCADE,
        related_name='brand'
    )
    name = models.CharField(max_length=255)

    class Meta:
        ordering = ['name']
        unique_together = ('vehicle_type', 'name')

    def __str__(self) -> str:
        return f"{self.name} ({self.vehicle_type.name})"


class Model(models.Model):
    brand = models.ForeignKey(
        Brand,
        on_delete=models.CASCADE,
        related_name='model'
    )
    name = models.CharField(max_length=255)

    class Meta:
        unique_together = ('brand', 'name')

    def __str__(self) -> str:
        return self.name

    class Meta:
        ordering = ['name']


class Vehicle(models.Model):
    name = models.CharField(max_length=255,
                            null=True,
                            blank=True
                            )
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='vehicle'
    )
    model = models.ForeignKey(
        Model,
        on_delete=models.CASCADE,
        related_name='vehicle'
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
            MaxValueValidator(99999)
        ])
    insurance_date = models.DateField(null=True, blank=True)

    def __str__(self) -> str:
        if self.name:
            return f"{self.name}"
        return f"Model: {self.model.name}, Plate: {self.plate_number}"

    class Meta:
        ordering = ['name']
