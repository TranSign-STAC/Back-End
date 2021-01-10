package main

import (
	"context"
	"fmt"
	"net"

	pb "transign/gen"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTextToSignLangServer
}

func (s *server) Translate(ctx context.Context, in *pb.TextToSignLangRequest) (*pb.TextToSignLangResponse, error) {
	// save to db
	// forward to translation server
	// forwared response from translation server
	logrus.WithFields(logrus.Fields{
		"UUID": in.Uuid,
		"Text": in.Text,
	}).Info("Request")
	response := &pb.TextToSignLangResponse{RenderUrl: fmt.Sprintf("api.transign.io/video/%s", in.Text)}
	return response, nil
}

func main() {
	const PORT = 8000

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	defer conn.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterTextToSignLangServer(s, &server{})

	fmt.Printf("Server Online on: %d\n", PORT)

	if err := s.Serve(conn); err != nil {
		fmt.Println("Failed to serve: ", err)
	}
}
