default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o hello ./go/src/github.com/karolhor/go-workshops-challange/server

builddocker:
	docker build -t karolhor/build-go-workshops-challange -f ./Dockerfile.build .
	docker run -t karolhor/build-go-workshops-challange /bin/true
	mkdir -p ./server/bin
	docker cp `docker ps -q -n=1`:/hello ./server/bin/hello
	chmod 755 ./server/bin/hello


