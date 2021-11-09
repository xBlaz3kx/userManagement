FROM golang:1.17 as base
WORKDIR /user-management/src
COPY . .
RUN ls -lart

FROM base as dev
ENTRYPOINT ["go","run","."]

FROM base as run

RUN go mod download
RUN go build -o app .
ENTRYPOINT ["app"]

FROM base as test
RUN go test -v

FROM alpine as app

COPY --from=run /user-management/src/app /usr/bin/app
COPY ./config /user-management/config

ENTRYPOINT ["app"]