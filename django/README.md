# Monolithic approach 

- [Monolithic approach](#monolithic-approach)
  - [TODOs](#todos)
  - [Requirements](#requirements)
  - [Commands](#commands)
  - [Routes](#routes)
  - [Services](#services)
  - [Notes](#notes)
## TODOs

If having more time, it is important to:
 - Create more unit tests (mostly on the sync related tasks) 
 - Create e2e tests 
 - Improve the error handling over Huey tasks
 - Add a policy for third-party data, since it relies on Ghibli's data integrity (for example: Ghibli's re-addition of movies would modify its ids, and then create a duplicated record on our side).
## Requirements

- docker-compose

## Commands

- First time setup & run: `./setup`

- Only worker (as daemon): `./run_worker`
> It'll create a worker.log in the root, with stdout and stderr of the subprocess

- Run: `./run`
> Keep in mind that `run` wraps an execution of the workers.

- Run tests: `./run_tests` (additionally, you can pass opt args/flags used in `manage.py test` directly, like this: `./run_tests --verbosity=2`)


## Routes

- `localhost:8000/movies`
- `localhost:8000/movies/.:movie-slug`

## Services

The project contains the following services (declared in the compose yaml):

- app: the main application service. It serves the Django project along with all its apps.
- db: the Postgres service. It serves a persistent storage for the app.
- redis-huey: redis store to use HueyRedis. It serves a redis server to workers (in app), which will perform tasks.

## Notes

I considered using Huey as the best approach for the project since I think Celery or Django Background Tasks are an overhead for that minimal use case.

Its a need to improve style code (I don't have a good Python environment due to I'm not on my laptop).
