package cmd

import (
	"fmt"
	"github.com/jkdv-systeme/kyasshu/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "dynamail",
	Short: "Transactional and marketing email service",
	Long:  `simply send emails with dynamail, track opens and clicks and much more`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Printf("dynamail - v%s (%s)\n", config.Version, config.Date)
		}
		_ = cmd.Help()
	},
}

func Execute() {
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "version")

	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Failed to execute command")
		os.Exit(1)
	}
}
