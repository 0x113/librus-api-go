FROM golang:1.14.4

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o librus-api-http api/main.go

WORKDIR /app

RUN cp /build/librus-api-http .

EXPOSE 8000
ENTRYPOINT ["/app/librus-api-http"]
