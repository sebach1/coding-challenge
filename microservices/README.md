# Microservices approach

- [Microservices approach](#microservices-approach)
  - [TODOs](#todos)
  - [Requirements](#requirements)
  - [Commands](#commands)
  - [Routes](#routes)
  - [Structure](#structure)
    - [Backend](#backend)
    - [Frontend](#frontend)
    - [Pb](#pb)

## TODOs

If having more time, it is important to:
 - Perform tests (using pb/pb<pkg>test mocks as long as relying in a gRPC svc)
 - Isolate people svc from films one if wanted
 - Add /movies/:slug route serving the movie's detail 
 - Add real-time updates on frontend using the gRPC bidirectional streaming over HTTP/2


## Requirements

- docker-compose

## Commands

- First time setup & run: `./setup` 

- Run: `./up`

## Routes

- `localhost:8000/movies`

## Structure 

### Backend

- ./backend/films
   
    a. films (`:9990`) 

    Responsible for serving all CRUD operations over films, people and their relations. 

    b. proxy (`:10000`) 

    Using grpc-go-gateway proxies all incoming requests from browser to gRPC API (listening at :9990) and transforms the protobuf response to a jsonified version to provide REST-like API.


- ./backend/ghibli

The service is a scheduler with a jobs handler, which manages tasks in a lightweight-way (without Redis or crontab). It's responsible for syncing with the Ghibli API using the films svc.

### Frontend

- ./frontend: holds the frontend (exposed at `:8000` as a Vue Single Web App).

### Pb

- pb/pbconf: svc configuration module.
- pb/<pkg>/*.pb.go: grpc generated code for golang client and server.
- pb/def/<pkg>.proto: proto declarations.