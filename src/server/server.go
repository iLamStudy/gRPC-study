package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/freewebsys/grpc-go-demo/src/helloworld"
	"google.golang.org/grpc/reflection"
	"fmt"
)

const (
	port = ":50051"
)

func (s *server) SayHello (ctx context.Context , in *pb.HelloRequest) (*pb.HelloReply , error) {
	fmt.Println("######## get Client request name :" + in.name)

	return &pb.HelloReply { Message:"Hello " + in.Name },nil
}

func main () {
	lis,err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to listen : %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer (s,&server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v" , err)
	}
}
