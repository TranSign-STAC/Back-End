package main

import (
	"fmt"
	"net"

	"transign/cmd/server/config"
	"transign/cmd/server/controllers"
	pb "transign/gen"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	config.ConnectDB()
	// connect db

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Envs["SERVER_PORT"]))
	defer conn.Close()
	if err != nil {
		logrus.Fatal(err)
		panic("Failed to listen.")
	}

	s := grpc.NewServer()

	pb.RegisterTextToSignLangServer(s, &controllers.TextToSignLangServer{})
	pb.RegisterTranslationHistoryServer(s, &controllers.TranslationHistoryServer{})
	// register controller functions

	logrus.Info(fmt.Sprintf("Server Online on: %s\n", config.Envs["SERVER_PORT"]))

	if err := s.Serve(conn); err != nil {
		logrus.Fatal(err)
		panic("Failed to serve.")
	}
}
