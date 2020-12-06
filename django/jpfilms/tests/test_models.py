from jpfilms.models import Film, People
from django.test import TestCase

class FilmModelTests(TestCase):

    def test_to_str(self):
        title = "foo"
        film = Film(title=title)
        self.assertIs(film.__str__(), title)


class PeopleModelTests(TestCase):

    def test_to_str(self):
        name = "foo"
        ppl = People(name=name)
        self.assertIs(ppl.__str__(), name)


