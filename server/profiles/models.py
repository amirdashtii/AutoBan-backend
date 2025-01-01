from django.db import models
from django.conf import settings
from django.utils.translation import gettext_lazy as _

from autoban.common.models import BaseModel


class Profile(BaseModel):
    user = models.OneToOneField(
        settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    first_name = models.CharField(_("first name"), max_length=255, blank=True)
    last_name = models.CharField(_("last name"), max_length=255, blank=True)
    birth_date = models.DateField(null=True, blank=True)
    GENDER_CHOICES = [
        ('M', 'Male'),
        ('F', 'Female'),
        ('O', 'Other'),
    ]
    gender = models.CharField(
        max_length=1, choices=GENDER_CHOICES, null=True, blank=True)

    def __str__(self):
        return f"{self.first_name} - {self.last_name}"
