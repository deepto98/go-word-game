FROM golang:alpine as builder

WORKDIR /go/src/app
ENV GO111MODULE=on

# For hot reload
RUN go get github.com/cespare/reflex@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

# Production:
FROM alpine:latest
RUN apk --nocache add ca-certificates
WORKDIR /root/

# Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8000
CMD ["./run"]