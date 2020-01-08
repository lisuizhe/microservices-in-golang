.PHONY: proto build clean run all

all: build clean run

proto:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto

build: proto
	RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	docker build -t shippy-service-consignment .

run: 
	docker run -p 50051:50051 \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  shippy-service-consignment

clean:
	rm consignment-service