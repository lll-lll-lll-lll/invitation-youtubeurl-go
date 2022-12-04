FROM golang:1.18-alpine as build-reflex
RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    rm -rf /var/cache/apk/*

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add bash git build-base jq && \
    rm -rf /var/cache/apk/*

COPY go.mod .
COPY go.sum .
RUN go mod download
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
COPY ./ ./
RUN go get github.com/lib/pq
RUN go build cmd/main.go
EXPOSE 8080
CMD [ "./main" ]