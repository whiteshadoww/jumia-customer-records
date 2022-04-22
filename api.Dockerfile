FROM golang:alpine as builder


RUN apk add build-base
WORKDIR /go/src/jumia
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . /go/src/jumia
RUN go build -o ./dist/jumia -buildvcs=false

FROM alpine:3.11.3
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Africa/Nairobi /etc/localtime && \
  apk del tzdata

COPY ./config/config.yaml .
COPY ./sample.db .
COPY --from=builder /go/src/jumia/dist/jumia .
EXPOSE 8888
ENTRYPOINT ["./jumia"]
