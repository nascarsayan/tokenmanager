/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"context"
	"fmt"
	"time"

	pb "example.com/tokenmanager/pkg/token"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReadOptions struct {
	ClientOptions
}

func newCmdRead(co *ClientOptions) *cobra.Command {
	opts := ReadOptions{
		ClientOptions: *co,
	}

	cmd := &cobra.Command{
		Use:   "read",
		Short: "Read token for client",
		Run: func(cmd *cobra.Command, args []string) {
			readToken(&opts)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateCommonFlags(&opts.ClientOptions)
		},
	}
	cmd = getCommonFlags(cmd, &opts.ClientOptions)

	return cmd
}

func readToken(opts *ReadOptions) {
	log := opts.getLogger()
	log.Info("Reading Token")
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%s", opts.Hostname, opts.Port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.ReadTokenRequest{
		Id: opts.ID,
	}
	r, err := client.ReadToken(ctx, &req)
	if err != nil {
		log.Fatalf("could not read token: %v", err)
	}
	fmt.Println(r.String())
}
