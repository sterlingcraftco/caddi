package cmd

import (
	"fmt"

	"github.com/sterlingcraftco/caddi/internal/caddy"

	"github.com/spf13/cobra"
)

var (
	newPort     string
	newUpstream string
)

var modifyCmd = &cobra.Command{
	Use:   "modify [server_name]",
	Short: "Modify a Caddy config block",
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
			fmt.Println("server not found")
			return
		}

		if newPort != "" {
			server.Listen = []string{newPort}
			fmt.Println("port updated to", newPort)
		}

		// optional: modify upstreams
		if newUpstream != "" {
			for i := range server.Routes {
				if len(server.Routes[i].Handle) > 0 {
					server.Routes[i].Handle[0].Upstreams = []caddy.Upstream{{Dial: newUpstream}}
				}
			}
			fmt.Println("upstream updated to", newUpstream)
		}

		config.Apps.HTTP.Servers[serverName] = server
		err = caddy.SaveConfig(config)
		if err != nil {
			fmt.Println("failed to save config:", err)
			return
		}
		fmt.Println("modification saved")
	},
}

func init() {
	modifyCmd.Flags().StringVarP(&newPort, "port", "p", "", "New port to listen on")
	modifyCmd.Flags().StringVar(&newUpstream, "upstream", "", "New reverse proxy upstream address")
	rootCmd.AddCommand(modifyCmd)
}
