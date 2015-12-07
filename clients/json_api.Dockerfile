FROM tianon/true

ADD ./bin/json_api .
ADD ./config/docker/json_api_client.json .

WORKDIR /
CMD ["./json_api", "--config", "json_api_client.json"]

EXPOSE 8080