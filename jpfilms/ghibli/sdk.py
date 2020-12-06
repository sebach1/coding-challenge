import requests
from jpfilms.models import Film, People

base_url = "https://ghibliapi.herokuapp.com"


def get_films(
    limit: int = 20,
    fields: [str] = [
        "id", "title", "description", "director", "producer", "release_date", "rt_score",
    ],
) -> dict:
    params = {"limit": limit, "fields": ",".join(fields)}
    resp = requests.get(base_url + "/films", params=params)
    return resp.json()


def get_people(
    limit: int = 20,
    fields: [str] = ["id", "name", "gender", "age", "eye_color", "hair_color", "films"],
) -> dict:
    params = {"limit": limit, "fields": ",".join(fields)}
    resp = requests.get(base_url + "/people", params=params)
    return resp.json()