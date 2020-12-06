CREATE TABLE IF NOT EXISTS films(
    id serial primary key,
    title varchar(60) not null,
    slug varchar(70) unique not null,
    rating integer,
    external_reference varchar(255) unique not null,
    director_name varchar(100) not null,
    producer_name varchar(100) not null,
    description text,
    release_year integer
);