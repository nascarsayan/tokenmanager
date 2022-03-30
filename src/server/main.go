package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc.io/tokenman/src/pb"
)

func main() {

	var port string

	flag.StringVar(&port, "port", "", "Port of the server")
	flag.Parse()

	if port == "" {
		log.Fatalf("port not provided")
	}

	addr := port
	if addr[0] != ":"[0] {
		addr = ":" + addr
	}

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("start server failed: %s\n", err)
	}
	log.Printf("server started at address: %s", addr)
	grpcServer := grpc.NewServer()

	pb.RegisterTokenServer(grpcServer, &Server{})

	log.Printf("GRPC server listening on %v", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
