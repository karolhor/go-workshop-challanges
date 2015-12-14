# go-workshops-challange
Challenge result for https://github.com/exu/go-workshops/blob/master/CHALLANGE.md

You can run whole client/server apps in 2 modes

1. complex - all apps and dependencies are in docker environment
2. local - only deps (mongo, redis) are run in docker 

## 1. Complex Docker Environment Architecture

For this challenge I have two separate containers environment

1. build container for compiling go services
2. docker-compose env with all dependencies (e.g. redis) & minimal containers for server/clients (1 per service)
 * **redis** - image container
 * **server** - ~7 MB container
 * **json_api** - ~6 MB container
 * **logger** - 
 * **mongo** -
 * **event_stream** -

### Run
```bash
$ make       # this will compile go services in docker container
$ make run   # this will run docker-compose environment 
```

## 2. Local mode


