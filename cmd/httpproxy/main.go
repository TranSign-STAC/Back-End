package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "transign/gen"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8000", "gRPC server endpoint")
)

const PORT = 8080

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterTextToSignLangHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := gw.RegisterTranslationHistoryHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := gw.RegisterFavoriteTranslationHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	// register handlers

	fmt.Printf("Server Online on: %d\n", PORT)
	return http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
