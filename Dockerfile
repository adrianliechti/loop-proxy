FROM golang:1-alpine3.16 as build

WORKDIR /src

COPY main.go .

RUN go build -o proxy main.go


FROM alpine:3.16

WORKDIR /app
COPY --from=build /src/proxy .

EXPOSE 8080

ENTRYPOINT [ "/app/proxy" ]