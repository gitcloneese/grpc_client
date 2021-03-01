package main

import (
	"context"
	"fmt"

	pd "github.com/gitcloneese/testgrpc"
	"google.golang.org/grpc"
)

func main() {
	// 客服端链接服务器
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())

	if err != nil {
		fmt.Println("网络异常", err)
	}

	defer conn.Close()

	c := pd.NewHelloserverClient(conn)

	re, err := c.Sayhello(context.Background(), &pd.HelloReq{Name: "熊猫"})

	if err != nil {
		fmt.Println("Sayhello 服务器调用失败！")
	}

	fmt.Println("调用Sayhello 服务器返回： ", re.Msg)

	/// 调用SayName

	re1, err := c.Sayname(context.Background(), &pd.NameReq{Name: "平平"})

	if err != nil {
		fmt.Println("Sayname 服务器调用失败 ！")
	}

	fmt.Println("调用Sayname 服务器返回： ", re1.Msg)

}
