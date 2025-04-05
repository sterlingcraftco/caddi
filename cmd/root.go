package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caddi",
	Short: "Caddi is a CLI tool to manage Caddy server configs",
	Long:  `Caddi lets you list, view, and modify Caddy server configuration blocks.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
