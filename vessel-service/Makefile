.PHONY: proto build clean run all

all: build clean run

proto:
	protoc -I. --go_out=plugins=micro:. \
	  proto/vessel/vessel.proto

build: proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-vessel .

run: 
	docker run -p 50052:50051 \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  shippy-service-vessel

clean:
	rm vessel-service