module github.com/359025765/go-micro-demo/demo-service

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/micro/go-micro v1.18.0
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20200930005306-bb64fee312b4 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
