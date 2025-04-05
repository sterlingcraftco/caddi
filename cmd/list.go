package cmd

import (
	"fmt"

	"github.com/sterlingcraftco/caddi/internal/caddy"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Caddy configuration blocks",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := caddy.LoadConfig()
		if err != nil {
			fmt.Println("error loading config:", err)
			return
		}
		i := 1
		for name, site := range config.Apps.HTTP.Servers {
			fmt.Printf("[%d] %s -> %+v\n", i, name, site.Listen)
			i++
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
