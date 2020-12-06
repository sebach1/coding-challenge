1. No hay criterio sobre si el 3rd party almacena bien la data (ej: restore de ids)
2. Mejorar err handling

# Django project

## Requirements

- docker-compose

## Commands 

- First time setup & run: `./setup` (requires docker-compose).

- Run: `./run`

- Only worker: `./run_worker`

Keep in mind that `run` wraps an execution of the workers.

## Routes

 - `localhost:8000/movies`

## Services 

The project contains the following services:
  - app: the main application service. It serves the Django project along with all its apps.
  - db: the postgres service. It serves a persistent storage for the app.
  - redis-huey: redis store to use HueyRedis. It serves a redis server to workers (in app), which will perform tasks.

