FROM ubuntu:16.04

RUN rm -rf /app
COPY main /app/mworker
WORKDIR /app

ENTRYPOINT ["/app/mworker"]