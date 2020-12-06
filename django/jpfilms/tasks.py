import time
from django.contrib.admin.utils import flatten
from huey import crontab
from huey.contrib import djhuey as huey
from jpfilms.models import Film, People
from jpfilms.ghibli.sdk import get_films, get_people
from jpfilms.ghibli import mappers
from jpfilms.serializers import people_from_dict, film_from_dict, create_slug 


@huey.db_periodic_task(crontab(minute='*/1'))
@huey.lock_task('ghibli-films-lock')
def sync_ghibli_films():
    current_quantity = Film.objects.count()
    remote_films = get_films(limit=250)
    if current_quantity == len(remote_films):
        return

    new_films = []
    for f in remote_films:
        film = film_from_dict(f, mapper=mappers.film)
        film.slug = create_slug(film.title, namespace="ghibli") 
        new_films.append(film)

    # Ignores conflicts to avoid duplication
    # of objects (see unique constraint over external_references)
    Film.objects.bulk_create(new_films, ignore_conflicts=True)

    return sync_ghibli_people()


def sync_ghibli_people():
    remote_people_list = get_people(limit=250)

    new_people = [] 
    for ppl in remote_people_list:
        people = people_from_dict(ppl, mapper=mappers.people)
        people.slug = create_slug(people.name, namespace="ghibli") 
        new_people.append(people)

    # Ignores conflicts to avoid duplication
    # of objects (see unique constraint over external_references)
    People.objects.bulk_create(new_people, ignore_conflicts=True)

    return sync_ghibli_relations(remote_people_list)


def sync_ghibli_relations(remote_people_list):
    Through = People.films.through
    throughs = []

    current_film_refs = Film.objects.values_list("id", mappers.film["id"])
    current_people_refs = People.objects.values_list(
        "id", mappers.people["id"])

    for ppl in remote_people_list:
        people_id = next(
            pid for pid, ext_ref in current_people_refs if ext_ref == ppl["id"])

        def get_film_id(film_url): return film_url.rsplit('/', 1)[-1]
        remote_film_ids = [get_film_id(film_url)
                           for film_url in ppl.pop("films")]
        film_refs = list(
            filter(lambda ref: ref[-1] in remote_film_ids, current_film_refs)
        )
        for film_id, _ in film_refs:
            throughs.append(
                Through(film_id=film_id, people_id=people_id)
            )
    Through.objects.bulk_create(throughs, ignore_conflicts=True)
    return
