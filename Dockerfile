FROM golang:1.18.0-alpine3.14 as builder
WORKDIR /go/src/github.com/megamsquare/dsanda
COPY . .
RUN go clean -modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/github.com/megamsquare/dsanda/cmd/main.go

FROM alpine:3.14
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/megamsquare/dsanda/main /app/main
EXPOSE 9010 9010
ENTRYPOINT ["/app/main"]


