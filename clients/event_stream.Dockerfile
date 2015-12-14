FROM tianon/true

ADD ./bin/event_stream .
ADD ./config/docker/event_stream.json .
ADD ./event_stream/static static/

WORKDIR /
CMD ["./event_stream", "--config", "event_stream.json"]

EXPOSE 8080