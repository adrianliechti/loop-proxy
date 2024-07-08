FROM golang:1-alpine AS build

WORKDIR /src

COPY main.go .

RUN go build -o proxy main.go


FROM alpine:3

WORKDIR /app
COPY --from=build /src/proxy .

EXPOSE 8080

ENTRYPOINT [ "/app/proxy" ]