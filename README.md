# go-workshops-challange
Challenge result for https://github.com/exu/go-workshops/blob/master/CHALLANGE.md

You can run whole client/server apps in 2 modes

1. complex - all apps and dependencies are in docker environment
2. local - only deps (mongo, redis) are run in docker 

## 1. Complex Docker Environment Architecture

This is a combination of two separate containers environment

1. build container for compiling go services
2. docker-compose env with all dependencies (e.g. redis) & minimal containers for server/clients (1 per service)
 * **redis** - image container
 * **server** - ~8 MB container
 * **json_api** - ~7.5 MB container
 * **logger** - ~5 MB container
 * **mongo** - ~6 MB container
 * **event_stream** - ~9 MB container

### Run
```bash
$ make       # this will compile go services in docker container
$ make run   # this will run docker-compose environment 
```

or simpler (with only one command)

```bash
$ make complex
```

Default binded ports:

* 8080 - server port
* 8888 - event_stream client port

## 2. Local mode

You can also run all clients/server app locally and keep only deps in docker env.

### Run

```bash
$ make docker-up-minimal  # runs redis & mongo
$ go run app.go           # runs all clients/server
```

Default binded ports:

* 8080 - server port
* 8081 - json_api client port
* 8082 - event_stream client port

# Server API

POST / 

{
    "message" : "Example"
}