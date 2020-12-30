FROM golang:alpine AS builder
RUN apk add --no-cache git make bash
RUN mkdir /go/src/go-skeleton
WORKDIR /go/src/go-skeleton
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN apk --no-cache add tzdata
ENV TZ Asia/Jakarta
EXPOSE 8080
COPY --from=builder /go/src/go-skeleton/bin/go-skeleton /
COPY --from=builder /go/src/go-skeleton/files/etc/skeleton /
ENTRYPOINT ["/go-skeleton"]
