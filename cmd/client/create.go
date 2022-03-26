/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
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
			fmt.Println("create called with options")
			createToken(&opts)
		},
	}

	return cmd
}

func createToken(opts *CreateOptions) {
	logger := opts.getLogger()
	logger.Info("Creating Token")
}
