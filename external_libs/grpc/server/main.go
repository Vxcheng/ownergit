package main

import (
	"context"
	"log"
	"net"

	"ownergit/external_libs/grpc/pb" // 添加了缺失的导入路径

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) StreamNumbers(in *pb.NumberRequest, stream pb.Greeter_StreamNumbersServer) error {
	for i := 1; i <= int(in.Max); i++ {
		if err := stream.Send(&pb.NumberReply{Number: int32(i)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
