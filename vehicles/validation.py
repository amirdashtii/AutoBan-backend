from django.core.exceptions import ValidationError
from django.core.validators import RegexValidator


class VehiclePlateValidator:
    car_plate_validator = RegexValidator(
        regex=r'^\d{2}[بپتثجچحخدذرزسشصطعفقکگلمنوهی]\d{3}\d{2}$')
    motorcycle_plate_validator = RegexValidator(regex=r'^\d{8}$')

    @staticmethod
    def validate(value):
        try:
            VehiclePlateValidator.car_plate_validator(value)
        except ValidationError:
            try:
                VehiclePlateValidator.motorcycle_plate_validator(value)
            except ValidationError:
                raise ValidationError("Enter valid plate")
