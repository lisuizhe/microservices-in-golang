FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]

# FROM golang:1.9.0 as builder

# WORKDIR /go/src/github.com/lisuizhe/microservices-in-golang/consignment-service
# COPY . .

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init && dep ensure
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# RUN mkdir /app
# WORKDIR /app
# COPY --from=builder /go/src/github.com/lisuizhe/microservices-in-golang/consignment-service .

# CMD ["./consignment-service"]