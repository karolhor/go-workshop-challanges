FROM tianon/true

ADD ./bin/logger .
ADD ./config/docker/logger.json .

WORKDIR /
CMD ["./logger", "--config", "logger.json"]