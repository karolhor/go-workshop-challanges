default: build-docker

app_dir = $$GOPATH/src/github.com/karolhor/go-workshops-challange

buildgo:
	rm -rf $(app_dir)/*
	cp -r /app/* $(app_dir)
	cd $(app_dir) && go get -v ./...

	cd $(app_dir)/server && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /server .
	cd $(app_dir)/clients/json_api && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /json_api .
	cd $(app_dir)/clients/logger && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /logger .
	cd $(app_dir)/clients/mongo && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /mongo .



build-on-docker:
	docker-compose -f docker-compose.build.yml build
	docker-compose -f docker-compose.build.yml up

build_container = $(shell docker ps -a | grep goworkshopschallange_build_1 | head -n 1 |  awk '{print $$1}')

copy-bin-docker:
	mkdir -p ./server/bin
	mkdir -p ./clients/bin

	docker cp $(build_container):/server ./server/bin/server
	docker cp $(build_container):/json_api ./clients/bin/json_api
	docker cp $(build_container):/logger ./clients/bin/logger
	docker cp $(build_container):/mongo ./clients/bin/mongo

	chmod 755 ./server/bin/server
	chmod 755 ./clients/bin/*

build-docker: build-on-docker copy-bin-docker

run:
	docker-compose build
	docker-compose up

docker-up-minimal:
	docker-compose -f docker-compose.yml -f docker-compose.minimal.yml up