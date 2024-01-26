FROM ubuntu:noble

COPY --link . /app
WORKDIR /app
RUN chmod +x ./sharks

CMD [ "./sharks" ]
