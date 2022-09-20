# syntax=docker/dockerfile:1

FROM golang:1.18.0-alpine3.14 as builder
COPY go.mod go.sum /go/src/github.com/megamsquare/dsanda/
WORKDIR /go/src/github.com/megamsquare/dsanda
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.14
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/megamsquare/dsanda/main /app/
EXPOSE 8080 8080
ENTRYPOINT ["/app/main"]


