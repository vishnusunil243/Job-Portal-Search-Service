FROM golang:1.21.5-bullseye AS build

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o search-service

FROM busybox:latest

WORKDIR /search/cmd

COPY --from=build /app/cmd .

COPY --from=build /app/.env /search

EXPOSE 8083

CMD ["./search-service"]