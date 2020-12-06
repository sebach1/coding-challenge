from jpfilms.serializers import SIG_IGNORE_KEY

film = {
    "id": "external_reference",
    "director": "director_name",
    "producer": "producer_name",
    "release_date": "release_year",
    "rt_score": "rating",
}

people = {
    "id": "external_reference",
    "films": SIG_IGNORE_KEY,
}
