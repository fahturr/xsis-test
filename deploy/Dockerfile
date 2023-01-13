FROM golang:1.19-alpine as build

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/

# Insert binary to small apline linux
FROM alpine

WORKDIR /app

COPY --from=build /app/main .

# run the service on container startup.
CMD ["/app/main"]