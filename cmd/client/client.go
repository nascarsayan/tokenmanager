/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"example.com/tokenmanager/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type ClientOptions struct {
	Hostname string
	Port     string
	ID       string
}

func NewCmdClient() *cobra.Command {
	var opts ClientOptions

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Manage tokens for client",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Hostname == "" {
				return cmdutil.FlagErrorf("--hostname cannot be empty")
			}
			if opts.Port == "" {
				return cmdutil.FlagErrorf("--port cannot be empty")
			}
			if opts.ID == "" {
				return cmdutil.FlagErrorf("--id cannot be empty")
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&opts.Hostname, "host", "", "The hostname of the server.")
	cmd.PersistentFlags().StringVar(&opts.Port, "port", "", "The port on which the server is running.")
	cmd.PersistentFlags().StringVar(&opts.ID, "id", "", "The id of the token.")

	cmd.AddCommand(newCmdCreate(&opts))
	cmd.AddCommand(newCmdDrop(&opts))
	cmd.AddCommand(newCmdWrite(&opts))
	cmd.AddCommand(newCmdRead(&opts))
	return cmd
}
