import uuid

from django.contrib.auth.models import AbstractUser
from django.contrib.auth.validators import UnicodeUsernameValidator
from django.db import models
from django.utils.translation import gettext_lazy as _
from django.core.validators import RegexValidator


class User(AbstractUser):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    phone_number = models.CharField(
        max_length=15,
        unique=True,
        help_text=_(
            "Phone number must be entered in the format: '09999999999' or '+999999999999'."),
        validators=[
            RegexValidator(
                regex=r'^\+?\d{10,15}$',
                message=_(
                    "Phone number must be entered in the format: '09999999999' or '+999999999999'.")
            ),
        ],
    )
    email = models.EmailField(
        _("email address"),
        unique=True,
        blank=True,
        null=True,
        help_text=_("Enter a valid email address."),
        error_messages={
            "unique": _("A user with that email address already exists."),
        },
    )
    username = models.CharField(
        _("username"),
        max_length=150,
        unique=True,
        blank=True,
        null=True,
        help_text=_(
            "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        ),
        validators=[UnicodeUsernameValidator()],
        error_messages={
            "unique": _("A user with that username already exists."),
        },
    )
    first_name = None
    last_name = None

    def is_staff(self):
        return self.is_admin

    USERNAME_FIELD = 'phone_number'
