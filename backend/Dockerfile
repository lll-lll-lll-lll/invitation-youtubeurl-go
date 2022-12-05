FROM golang:1.18-alpine as build-reflex
RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN apk update && \
    apk upgrade && \
    apk add bash git build-base jq && \
    rm -rf /var/cache/apk/*


ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
COPY backend/go.mod backend/go.sum ./
COPY ./ /go/src/app
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "cmd/main.go"]