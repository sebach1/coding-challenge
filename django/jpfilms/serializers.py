from .models import Film, People
from django.template.defaultfilters import slugify

SIG_IGNORE_KEY = False 

'''
map_dict performs a relative mapping given a mapper
for example:
    d: {"a": 1, "b": 2, "c": 3}
    mapper: {"a": "aa", "c": SIG_IGNORE_KEY}
    
    results in: {"aa": 1, "b": 2}
'''
def map_dict(d: dict, mapper: dict) -> dict:
    mapped_dict = {}
    for k, v in d.items():
        if k in mapper:
            if mapper[k] == SIG_IGNORE_KEY:
                continue
            k = mapper[k]
        mapped_dict[k] = v
    return mapped_dict


def create_slug(s: str, namespace: str) -> str:
    return slugify(namespace + " " + s) 

def film_from_dict(d: dict, mapper: dict) -> Film:
    return Film(**map_dict(d=d, mapper=mapper))


def people_from_dict(d: dict, mapper: dict) -> People:
    return People(**map_dict(d=d, mapper=mapper))
