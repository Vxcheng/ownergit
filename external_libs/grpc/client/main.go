package main

import (
	"context"
	"io"
	"log"
	"time"

	"ownergit/external_libs/grpc/pb" // 添加了缺失的导入路径

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 一元 RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	// 流式 RPC
	stream, err := c.StreamNumbers(context.Background(), &pb.NumberRequest{Max: 5})
	if err != nil {
		log.Fatalf("could not stream: %v", err)
	}
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream error: %v", err)
		}
		log.Printf("Received number: %d", num.Number)
	}
}
