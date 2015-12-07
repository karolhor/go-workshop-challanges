default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o server ./go/src/github.com/karolhor/go-workshops-challange/server
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o json_api ./go/src/github.com/karolhor/go-workshops-challange/clients/json_api
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o logger ./go/src/github.com/karolhor/go-workshops-challange/clients/logger

builddocker:
	docker build -t karolhor/build-go-workshops-challange -f ./Dockerfile.build .
	docker run -t karolhor/build-go-workshops-challange /bin/true
	mkdir -p ./server/bin
	mkdir -p ./clients/bin
	docker cp `docker ps -q -n=1`:/server ./server/bin/server
	docker cp `docker ps -q -n=1`:/json_api ./clients/bin/json_api
	docker cp `docker ps -q -n=1`:/logger ./clients/bin/logger
	chmod 755 ./server/bin/server
	chmod 755 ./clients/bin/*

run:
	docker-compose build
	docker-compose up

docker-up-minimal:
	docker-compose -f docker-compose.yml -f docker-compose.minimal.yml up

