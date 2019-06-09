package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"fmt"
	pb "github.com/nantaphop/go-grpc-example/location"
	"github.com/nantaphop/go-grpc-example/server"
)

func main(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatal(":10000 already in use")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLocationServer(grpcServer, server.NewLocationServer())
	grpcServer.Serve(lis)

}