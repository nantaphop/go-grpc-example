package server

import (
	"context"
	pb "github.com/nantaphop/go-grpc-example/location"
)
// LocationServer implementation
type LocationServer struct{}

// NewLocationServer return object that represent LocationServer
func NewLocationServer() *LocationServer {
	return new(LocationServer)
}

// Print location as string
func (s *LocationServer) Print(context.Context, *pb.LatLng) (*pb.StringResp, error){
	resp := &pb.StringResp{}
	resp.Message = "Hello"
	return resp, nil
}

