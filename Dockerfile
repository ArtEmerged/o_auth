FROM golang:1.22.6-alpine3.18 as builder
RUN apk update && apk add --no-cache git

WORKDIR /src
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -tags musl -a -installsuffix cgo -o app ./cmd/

FROM alpine:3.18
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /root/
RUN mkdir -p ./static
RUN apk update && apk add tzdata

ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /src/app .
CMD ["./app"]