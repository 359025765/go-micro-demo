package main

import (
	"context"
	pb "github.com/359025765/go-micro-demo/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"log"
)

type DemoServiceHandler struct {
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("laracom.demo.service"))
	service.Init()
	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
