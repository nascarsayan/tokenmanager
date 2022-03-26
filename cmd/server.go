/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ServerOptions struct {
	Hostname string
	Port     string
	ID       string
}

func NewCmdServer() *cobra.Command {
	var opts ServerOptions

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run a server for generating tokens",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called with options")
			fmt.Printf("%+v", opts)
		},
	}

	cmd.Flags().StringVar(&opts.Hostname, "host", "", "The hostname of the server.")
	cmd.Flags().StringVar(&opts.Port, "port", "", "The port on which the server is running.")
	cmd.Flags().StringVar(&opts.ID, "id", "", "The id of the token.")

	return cmd
}
