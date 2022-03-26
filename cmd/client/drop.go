/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
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
			fmt.Println("drop called with options")
			fmt.Printf("%+v", opts)
		},
	}

	return cmd
}
