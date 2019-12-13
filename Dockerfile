FROM golang:1.13.1-alpine3.10 AS build
WORKDIR /usr/src/WaterLogger-UserServer
COPY . .
RUN go build

FROM alpine:3.10 AS production
WORKDIR /app
COPY --from=build /usr/src/WaterLogger-UserServer/WaterLogger-UserServer .

EXPOSE 11011/tcp
ENTRYPOINT ["./WaterLogger-UserServer"]