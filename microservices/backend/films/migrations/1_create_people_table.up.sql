CREATE TABLE IF NOT EXISTS people(
    id serial primary key,
    name varchar(60) not null,
    gender varchar(30),
    age varchar(60),
    hair_color varchar(30),
    eye_color varchar(30),
    slug varchar(70) unique not null,
    external_reference varchar(255) unique not null
);