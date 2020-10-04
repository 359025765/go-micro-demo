module github.com/359025765/go-micro-demo/demo-cli

go 1.14

require (
	github.com/359025765/go-micro-demo/demo-service v0.0.0-20200930075126-d021b21c9a60
	github.com/micro/go-micro v1.18.0
    google.golang.org/grpc v1.31.0
)

replace github.com/359025765/go-micro-demo/demo-service => /Users/chenjiaxing/Workspace/go/laracom/demo-service

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
