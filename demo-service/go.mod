module github.com/359025765/go-micro-demo/demo-service

go 1.14

require (
	github.com/micro/go-micro v1.18.0
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
