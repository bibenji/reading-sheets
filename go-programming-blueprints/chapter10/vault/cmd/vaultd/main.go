package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"../../../vault"
	"../../pb"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8070", "http listen address")
		gRPCAddr = flag.String("grpc", ":8071", "gRPC listen address")
	)

	flag.Parse()

	ctx := context.Background()
	srv := vault.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}{}

	hashEndpoint := vault.MakeHashEndpoint(srv)
	validateEndpoint := vault.MakeValidateEndpoint(srv)
	endpoints := vault.Endpoints{
		hashEndpoint: hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("https:", *httpAddr)
		handler := vault.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := vault.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	// preventing the main function from terminating immediately
	log.Fatalln(<-errChan)
}
