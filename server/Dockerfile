FROM tianon/true

ADD ./bin/server .
ADD ./config/docker/server.json .

WORKDIR /
CMD ["./server", "--config", "server.json"]

EXPOSE 8080