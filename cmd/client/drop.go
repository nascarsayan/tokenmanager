/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DropOptions struct {
	Hostname string
	Port     string
	ID       string
}

func newCmdDrop(co *ClientOptions) *cobra.Command {
	opts := DropOptions{
		Hostname: co.Hostname,
		Port:     co.Port,
		ID:       co.ID,
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
