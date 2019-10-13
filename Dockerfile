FROM golang:1.13.1 AS builder

ADD . $GOPATH/github.com/the-redback/golang-api
WORKDIR $GOPATH/github.com/the-redback/golang-api

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /conways .

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /conways ./
RUN chmod +x ./conways

ENTRYPOINT ["./conways"]
EXPOSE 12345

