# Go
FROM golang:latest AS go_builder
WORKDIR /app
COPY /server ./
RUN go mod download && go mod verify
# RUN go build -o /main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# React
FROM node:alpine AS node_builder
COPY /client ./
RUN npm install
RUN npm run build

# Production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
ADD config /config
# COPY /server/config/ ./ ← NG
# ADD config /config ← TEST
# ADD asset /asset ← TEST
# ADD views /views ← TEST
COPY --from=go_builder /main ./
COPY --from=node_builder /build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main

#$ docker build -t ditto-go-react .
#$ docker run -p 80:8080 -d ditto-go-react