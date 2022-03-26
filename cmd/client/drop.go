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

type DropOptions struct {
	ClientOptions
}

func newCmdDrop(co *ClientOptions) *cobra.Command {
	opts := DropOptions{
		ClientOptions: *co,
	}

	cmd := &cobra.Command{
		Use:   "drop",
		Short: "Drop token for client",
		Run: func(cmd *cobra.Command, args []string) {
			dropToken(&opts)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateCommonFlags(&opts.ClientOptions)
		},
	}
	cmd = getCommonFlags(cmd, &opts.ClientOptions)

	return cmd
}

func dropToken(opts *DropOptions) {
	log := opts.getLogger()
	log.Info("Dropping Token")
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

	req := pb.DropTokenRequest{
		Id: opts.ID,
	}
	r, err := client.DropToken(ctx, &req)
	if err != nil {
		log.Fatalf("could not drop token: %v", err)
	}
	fmt.Println(r.String())
}
