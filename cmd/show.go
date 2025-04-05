package cmd

import (
	"fmt"

	"github.com/sterlingcraftco/caddi/internal/caddy"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [server_name]",
	Short: "Show details of a specific Caddy config block",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverName := args[0]
		config, err := caddy.LoadConfig()
		if err != nil {
			fmt.Println("error loading config:", err)
			return
		}
		server, ok := config.Apps.HTTP.Servers[serverName]
		if !ok {
			fmt.Printf("no server found with name: %s\n", serverName)
			return
		}
		fmt.Printf("Name: %s\n", serverName)
		fmt.Printf("Listen Ports: %v\n", server.Listen)
		fmt.Printf("Routes: %v\n", server.Routes)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
