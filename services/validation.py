from django.core.exceptions import ValidationError
from django.utils import timezone

def future_date_validator(value):
    if value > timezone.now().date():
        raise ValidationError('Service date cannot be in the future.')
