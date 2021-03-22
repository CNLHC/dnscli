package main

import (
	"fmt"
	"os"

	"github.com/CNLHC/dnscli/certbot"
	"github.com/CNLHC/dnscli/cmd"
	"github.com/CNLHC/dnscli/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "alidns",
		Short: "alidns cli",
		Run: func(c *cobra.Command, args []string) {
			fmt.Printf("Main Run")
		},
		PersistentPreRun: func(c *cobra.Command, args []string) {
			cfg := config.GetGlobalConfig()
			err := godotenv.Load(cfg.Dotfile)
			if err != nil {
				fmt.Printf("No configuration found")
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(certbot.GetCMD())
	cfg := config.GetGlobalConfig()
	rootCmd.PersistentFlags().StringVar(&cfg.Dotfile, "dotfile", ".env", "dotfile (default is .env)")
	cmd.DecorateRootCmd(rootCmd)

}

func main() {
	rootCmd.Execute()
}
