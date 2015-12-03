FROM tianon/true

ADD ./bin/json_api .

CMD ["/json_api"]

EXPOSE 8080