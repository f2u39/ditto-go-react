package rpc

import (
	"context"
	pb "ditto/rpc/act"
	"log"
	"net"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedActivityServer
}

func Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("failed to listen:", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterActivityServer(grpcSrv, &Server{})
	log.Println("gRPC server listening on", lis.Addr())
	if err := grpcSrv.Serve(lis); err != nil {
		log.Println("failed to serve:", err)
	}
}

func (s *Server) GetActs(ctx context.Context, req *pb.Req) (*pb.Rep, error) {
	res := &pb.Rep{}

	if req == nil {
		return res, xerrors.Errorf("request must not be nil")
	}

	return res, nil
}
