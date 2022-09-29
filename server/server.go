package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "internProject2/databus"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	oper     = flag.String("op", "div", "The operation used")
	operFunc func(float32, float32) (float32, string)
)

type server struct {
	pb.UnimplementedDatabusServiceServer
}

func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	prm1 := in.GetPrm1()
	prm2 := in.GetPrm2()
	res, str := operFunc(prm1, prm2)
	log.Println(str)
	return &pb.SendResponse{Result: res}, nil
}

func main() {
	var (
		operation    func(float32, float32) float32
		operationStr string
	)
	switch *oper {
	case "add":
		operation = func(p1, p2 float32) float32 { return p1 + p2 }
		operationStr = "+"
	case "sub":
		operation = func(p1, p2 float32) float32 { return p1 - p2 }
		operationStr = "-"
	case "mul":
		operation = func(p1, p2 float32) float32 { return p1 * p2 }
		operationStr = "*"
	case "div":
		operation = func(p1, p2 float32) float32 { return p1 / p2 }
		operationStr = "/"
	}

	operFunc = func(p1, p2 float32) (float32, string) {
		res := operation(p1, p2)
		str := fmt.Sprintf("%f %s %f = %f", p1, operationStr, p2, res)
		return res, str
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDatabusServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
