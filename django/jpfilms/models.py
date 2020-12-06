from django.db import models
from django.core import serializers

class Film(models.Model):
    title = models.CharField(max_length=60)
    description = models.TextField()
    release_year = models.IntegerField()
    director_name = models.CharField(max_length=100)
    producer_name = models.CharField(max_length=100)
    slug = models.CharField(max_length=70, unique=True, null=False)
    rating = models.IntegerField()
    external_reference = models.CharField(max_length=255, unique=True)

    def __str__(self):
        return '%s' % self.title


class People(models.Model):
    name = models.CharField(max_length=60)
    gender = models.CharField(max_length=30)
    age = models.CharField(max_length=60)  # Obs: age can be an age annotation, like "elder"
    hair_color = models.CharField(max_length=30)
    eye_color = models.CharField(max_length=30)
    films = models.ManyToManyField(Film)
    slug = models.CharField(unique=True, max_length=64, null=False)
    external_reference = models.CharField(max_length=255, unique=True)

    def __str__(self):
        return '%s' % self.name
