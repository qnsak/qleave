# build stage
FROM golang:alpine AS build-env
RUN apk add build-base
ADD . /src
RUN cd /src && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

# FROM  alpine
FROM alpine
ENV GIN_MODE=release
COPY ./.env /
COPY ./sqlitedb.db /
COPY --from=build-env /src/app /
EXPOSE 8000
WORKDIR /
ENTRYPOINT ["./app"]