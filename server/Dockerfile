FROM golang:1.19 AS go_builder

WORKDIR /usr/go/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o app

# $docker build -t ditto-go .
# CMD ["./app"]