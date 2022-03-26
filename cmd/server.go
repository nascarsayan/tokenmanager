/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"strings"

	"example.com/tokenmanager/pkg/cmdutil"
	"example.com/tokenmanager/pkg/server"
	pb "example.com/tokenmanager/pkg/token"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type ServerOptions struct {
	Port string
}

func (so *ServerOptions) getLogger() *log.Entry {
	return log.WithFields(log.Fields{
		"Port": so.Port,
	})
}

func NewCmdServer() *cobra.Command {
	var opts ServerOptions

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run a server for generating tokens",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Port == "" {
				return cmdutil.FlagErrorf("--port cannot be empty")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called with options")
			startServer(&opts)
		},
	}

	cmd.Flags().StringVar(&opts.Port, "port", "", "The port on which the server is running.")

	return cmd
}

func startServer(opts *ServerOptions) {
	logger := opts.getLogger()
	logger.Info("Starting server")

	address := opts.Port
	if !strings.HasPrefix(address, ":") {
		address = fmt.Sprintf(":%s", address)
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTokenServer(grpcServer, &server.Server{})

	log.Printf("GRPC server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Token struct {
	Id      string
	Name    string
	Low     uint64
	Mid     uint64
	High    uint64
	Partial uint64
	Final   uint64
}
