from rest_framework import serializers
from profiles.models import Profile


class ProfileSerializer(serializers.ModelSerializer):
    user_id = serializers.UUIDField(read_only=True)

    class Meta:
        model = Profile
        fields = ['id', 'user_id', 'first_name',
                  'last_name', 'birth_date', 'email', 'gender']
