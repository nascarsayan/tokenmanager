/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package client

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Hostname string
	Port     string
	ID       string
}

func newCmdCreate(co *ClientOptions) *cobra.Command {
	opts := CreateOptions{
		Hostname: co.Hostname,
		Port:     co.Port,
		ID:       co.ID,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create token for client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called with options")
			fmt.Printf("%+v", opts)
		},
	}

	return cmd
}
