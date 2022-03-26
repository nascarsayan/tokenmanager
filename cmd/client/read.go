/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
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
			fmt.Println("read called with options")
			fmt.Printf("%+v", opts)
		},
	}

	return cmd
}
