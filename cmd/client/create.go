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

type CreateOptions struct {
	ClientOptions
}

func newCmdCreate(co *ClientOptions) *cobra.Command {
	opts := CreateOptions{
		ClientOptions: *co,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create token for client",
		Run: func(cmd *cobra.Command, args []string) {
			createToken(&opts)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateCommonFlags(&opts.ClientOptions)
		},
	}
	cmd = getCommonFlags(cmd, &opts.ClientOptions)

	return cmd
}

func createToken(opts *CreateOptions) {
	log := opts.getLogger()
	log.Info("Creating Token")
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

	req := pb.CreateTokenRequest{
		Id: opts.ID,
	}
	r, err := client.CreateToken(ctx, &req)
	if err != nil {
		log.Fatalf("could not create token: %v", err)
	}
	fmt.Println(r.String())
}
