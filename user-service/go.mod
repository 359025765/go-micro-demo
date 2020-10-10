module github.com/359025765/go-micro-demo/user-service

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
