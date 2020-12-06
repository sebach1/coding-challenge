from jpfilms.models import Film, People
from django.test import TestCase
from jpfilms import serializers
from django.forms.models import model_to_dict


class SerializerTests(TestCase):

    def test_map_dict(self):
        d = {"a": 1, "b": 2, "c": 3}
        mapper = {"a": "aa", "c": serializers.SIG_IGNORE_KEY}
        expected = {"aa": 1, "b": 2}
        self.assertDictEqual(serializers.map_dict(
            d=d, mapper=mapper), expected)

    def test_people_from_dict(self):
        d = {"name": "foo", "aging": "bar", "films": [{"name": "baz"}]}
        mapper = {"aging": "age", "films": serializers.SIG_IGNORE_KEY}
        got = serializers.people_from_dict(d=d, mapper=mapper)
        expected = People(name=d["name"], age=d["aging"])
        self.assertDictEqual(model_to_dict(got), model_to_dict(expected))

    def test_film_from_dict(self):
        d = {"name": "bar", "people": [{"name": "foo"}]}
        mapper = {"name": "title", "people": serializers.SIG_IGNORE_KEY}
        got = serializers.film_from_dict(d=d, mapper=mapper)
        expected = Film(title=d["name"])
        self.assertDictEqual(model_to_dict(got), model_to_dict(expected))
