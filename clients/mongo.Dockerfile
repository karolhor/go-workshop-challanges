FROM tianon/true

ADD ./bin/mongo .
ADD ./config/docker/mongo.json .

WORKDIR /
CMD ["./mongo", "--config", "mongo.json"]