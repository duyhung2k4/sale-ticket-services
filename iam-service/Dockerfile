FROM golang:1.23-alpine AS build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /build/app /app/app
COPY --from=build /build/config/ /app/config/

ENTRYPOINT [ "/app/app" ]