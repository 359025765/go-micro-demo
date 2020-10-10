build:
	protoc --proto_path=. --micro_out=. --go_out=. proto/demo/demo.proto
	docker build -t demo-service -f .
run:
	docker run -d -p 9091:09091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns demo-service