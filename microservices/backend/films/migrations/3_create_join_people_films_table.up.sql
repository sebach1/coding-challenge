CREATE TABLE IF NOT EXISTS join_people_films(
    id serial primary key,
    film_id integer not null,
    people_id integer not null
);
ALTER TABLE ONLY join_people_films ADD CONSTRAINT uq_people_film_relation UNIQUE(film_id, people_id);