FROM golang:1.23 AS build

LABEL org.opencontainers.image.description DESCRIPTION

WORKDIR /app
COPY . /app
RUN go mod download && CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11:latest-amd64

COPY --from=build /go/bin/app /

ENV SERVICE_PORT 3000

EXPOSE ${SERVICE_PORT}

CMD ["/app"]