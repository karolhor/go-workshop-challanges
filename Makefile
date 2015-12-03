default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o hello ./go/src/github.com/karolhor/go-workshops-challange/server
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o json_api ./go/src/github.com/karolhor/go-workshops-challange/clients/json_api

builddocker:
	docker build -t karolhor/build-go-workshops-challange -f ./Dockerfile.build .
	docker run -t karolhor/build-go-workshops-challange /bin/true
	mkdir -p ./server/bin
	mkdir -p ./clients/bin
	docker cp `docker ps -q -n=1`:/hello ./server/bin/hello
	docker cp `docker ps -q -n=1`:/json_api ./clients/bin/json_api
	chmod 755 ./server/bin/hello
	chmod 755 ./clients/bin/json_api

run:
	docker-compose build
	docker-compose up
