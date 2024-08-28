
# Build the gobinary

FROM golang:1.22 AS gobuild

RUN update-ca-certificates

WORKDIR /go/src/app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN go mod verify

COPY ./ ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -linkmode external -extldflags -static" -o /go/bin/events-service ./cmds/events-service

FROM gcr.io/distroless/static

COPY --from=gobuild /go/bin/events-service /go/bin/events-service
EXPOSE 8090

ENTRYPOINT ["/go/bin/events-service"]
