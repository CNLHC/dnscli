package main

import (
	"os"

	"github.com/CNLHC/dnscli/certbot"
	"github.com/CNLHC/dnscli/cmd"
	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/pkg/ddns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config_path string
var (
	rootCmd = &cobra.Command{
		Use:   "dnscli",
		Short: "dnscli",
		Run: func(c *cobra.Command, args []string) {
			c.Help()
			os.Exit(0)
		},
		PersistentPreRun: func(c *cobra.Command, args []string) {
			if config_path != "" {
				viper.AddConfigPath(config_path)
			}
			if err := viper.ReadInConfig(); err != nil {
				logger := config.GetLogger()
				logger.Error().Msgf("no configuration")
				panic(err)

			}
		},
	}
)

func init() {

	viper.AddConfigPath(".")
	viper.AddConfigPath("~")
	viper.AddConfigPath(config.GetGoModRoot())
	viper.SetConfigName(".dnscli")
	viper.SetConfigType("yaml")

	rootCmd.PersistentFlags().StringVar(&config_path, "config", "~", "config file path")
	rootCmd.AddCommand(certbot.GetCMD())
	ddns.RegisterCMD(rootCmd)
	cmd.DecorateRootCmd(rootCmd)

}

func main() {
	rootCmd.Execute()
}
