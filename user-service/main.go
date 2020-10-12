package main

import (
	"fmt"
	database "github.com/359025765/go-micro-demo/user-service/db"
	"github.com/359025765/go-micro-demo/user-service/handler"
	pb "github.com/359025765/go-micro-demo/user-service/proto/user"
	repositroy "github.com/359025765/go-micro-demo/user-service/repo"
	"github.com/359025765/go-micro-demo/user-service/service"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	// 创建数据库连接，程序退出时断开连接
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 每次启动服务都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})

	// 初始化Repo实例用于后续数据库操作
	repo := &repositroy.UserRepository{db}

	// 初始化token service
	token := &service.TokenService{repo}

	// 创建微服务
	service := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
	)
	service.Init()

	// 注册处理器
	pb.RegisterUserServiceHandler(service.Server(), &handler.UserService{repo, token})

	// 启动用户服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
