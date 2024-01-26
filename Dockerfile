FROM ubuntu:noble

ADD . /app
WORKDIR /app
RUN chmod +x ./sharks

CMD [ "./sharks" ]
