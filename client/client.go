package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "internProject2/databus"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	prm1 = flag.Float64("prm1", 2, "first param")
	prm2 = flag.Float64("prm2", 3, "second param")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatabusServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.SendRequest{Prm1: float32(*prm1), Prm2: float32(*prm2)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.GetResult())
}
