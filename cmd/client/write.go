/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"context"
	"fmt"
	"time"

	"example.com/tokenmanager/pkg/cmdutil"
	pb "example.com/tokenmanager/pkg/token"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WriteOptions struct {
	ClientOptions
	Name string
	Low  uint64
	Mid  uint64
	High uint64
}

func newCmdWrite(co *ClientOptions) *cobra.Command {
	opts := WriteOptions{
		ClientOptions: *co,
	}

	cmd := &cobra.Command{
		Use:   "write",
		Short: "Write token for client",
		Run: func(cmd *cobra.Command, args []string) {
			writeToken(&opts)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCommonFlags(&opts.ClientOptions); err != nil {
				return err
			}
			if opts.Name == "" {
				return cmdutil.FlagErrorf("--name cannot be empty")
			}
			if !cmd.Flags().Changed("low") {
				return cmdutil.FlagErrorf("--low cannot be empty")
			}
			if !cmd.Flags().Changed("mid") {
				return cmdutil.FlagErrorf("--mid cannot be empty")
			}
			if !cmd.Flags().Changed("high") {
				return cmdutil.FlagErrorf("--high cannot be empty")
			}
			if !((opts.Low < opts.Mid) && (opts.Mid < opts.High)) {
				return cmdutil.FlagErrorf("The following condition is not met: low < mid < high")
			}
			return nil
		},
	}

	cmd = getCommonFlags(cmd, &opts.ClientOptions)
	cmd.Flags().StringVar(&opts.Name, "name", "", "The name of the token.")
	cmd.Flags().Uint64Var(&opts.Low, "low", 0, "The value low.")
	cmd.Flags().Uint64Var(&opts.Mid, "mid", 0, "The value low.")
	cmd.Flags().Uint64Var(&opts.High, "high", 0, "The value low.")
	return cmd
}

func writeToken(opts *WriteOptions) {
	log := opts.getLogger()
	log.Info("Writing Token")
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

	req := pb.WriteTokenRequest{
		Id:   opts.ID,
		Name: opts.Name,
		Low:  uint64(opts.Low),
		High: uint64(opts.High),
		Mid:  uint64(opts.Mid),
	}
	r, err := client.WriteToken(ctx, &req)
	if err != nil {
		log.Fatalf("could not write token: %v", err)
	}
	fmt.Println(r.String())
}
