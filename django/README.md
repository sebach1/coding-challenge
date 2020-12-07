1. No hay criterio sobre si el 3rd party almacena bien la data (ej: restore de ids)
2. Mejorar err handling

# Django project

## Requirements

- docker-compose

## Commands

- First time setup & run: `./setup` (requires docker-compose).

- Only worker (as daemon): `./run_worker`
> It'll create a worker.log in the root, with stdout and stderr of the subprocess

- Run: `./run`
> Keep in mind that `run` wraps an execution of the workers.

- Run tests: `./run_tests` (additionally, you can pass opt args/flags used in `manage.py test` directly, like this: `./run_tests --verbosity=2`)


## Routes

- `localhost:8000/movies`
- `localhost:8000/movies/.:movie-slug`

## Services

The project contains the following services:

- app: the main application service. It serves the Django project along with all its apps.
- db: the postgres service. It serves a persistent storage for the app.
- redis-huey: redis store to use HueyRedis. It serves a redis server to workers (in app), which will perform tasks.
