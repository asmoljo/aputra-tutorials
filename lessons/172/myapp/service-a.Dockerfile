FROM golang:1.20.6-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY cmd/service-a/main.go ./cmd/service-a/

RUN go build -o /myapp cmd/service-a/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=build /myapp /myapp

ENTRYPOINT ["/myapp"]