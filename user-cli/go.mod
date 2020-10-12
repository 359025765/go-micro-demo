module github.com/359025765/go-micro-demo/user-cli

go 1.14

require (
	github.com/359025765/go-micro-demo/user-service v0.0.0-20201010033258-ed00c31a32ac
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	golang.org/x/net v0.0.0-20201009032441-dbdefad45b89
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
