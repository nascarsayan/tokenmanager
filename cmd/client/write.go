/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
)

type WriteOptions struct {
	ClientOptions
	Name string
	Low  int
	Mid  int
	High int
}

func newCmdWrite(co *ClientOptions) *cobra.Command {
	opts := WriteOptions{
		ClientOptions: *co,
	}

	cmd := &cobra.Command{
		Use:   "write",
		Short: "Write token for client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("write called with options")
			fmt.Printf("%+v", opts)
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "The name of the token.")
	cmd.Flags().IntVar(&opts.Low, "low", -1, "The value low.")
	cmd.Flags().IntVar(&opts.Mid, "mid", -1, "The value low.")
	cmd.Flags().IntVar(&opts.High, "high", -1, "The value low.")
	return cmd
}
