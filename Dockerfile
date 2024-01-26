FROM ubuntu:noble

ADD . /app
WORKDIR /app

CMD [ "./sharks" ]