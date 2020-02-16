FROM golang:1.13-alpine

WORKDIR /go/src/go-web-server
COPY ./code .

RUN go get -d -v ./... \
    && go install -v ./... \
    && go build -o go-web-server ./cmd/web

CMD ["./go-web-server"]
