package main

import (
	pb "github.com/359025765/go-micro-demo/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {

	// 初始化客户端服务，定义命令行参数
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "test",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "test@qq.com",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "12345678",
			},
		),
	)

	// 远程服务客户端调用
	client := pb.NewUserService("laracom.user.service", service.Client())

	// 运行客户端命令调用远程服务逻辑
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 调用用户创建服务
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			log.Println("参数: ", name, email, password)
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
			})
			if err != nil {
				log.Fatalf("创建用户失败: %v", err)
			}
			log.Printf("创建用户成功: %s", r.User.Id)

			// 调用用户认证服务
			var token *pb.Token
			token, err = client.Auth(context.TODO(), &pb.User{
				Email:    email,
				Password: password,
			})
			if err != nil {
				log.Fatalf("用户登录失败: %v", err)
			}
			log.Printf("用户登录成功: %s", token.Token)

			// 调用用户验证服务
			token, err = client.ValidateToken(context.TODO(), token)
			if err != nil {
				log.Fatalf("用户认证失败: %v", err)
			}
			log.Printf("用户认证成功: %v", token.Valid)

			// 调用获取所有用户服务
			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("获取所有用户失败: %v", err)
			}
			log.Printf("所有用户: %s", getAll)
			/*for _, v := range getAll.Users {
				log.Println(v)
			}*/
			os.Exit(0)
		}),
	)
	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败: %v", err)
	}

}
