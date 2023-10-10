package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jkdv-systeme/kyasshu/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.KeyDelimiter("::")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/opt/kyasshu/")
	viper.AddConfigPath("/etc/kyasshu/")
	viper.AddConfigPath(".")

	viper.SetDefault("debug", false)

	viper.SetDefault("server.port", 4682)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("no config file found")
			os.Exit(1)
		} else {
			fmt.Println("failed to read config file")
			os.Exit(1)
		}
	}

	debugMode := viper.GetBool("debug")

	if debugMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	} else {
		zerolog.TimestampFieldName = "ts"
		zerolog.LevelFieldName = "level"
		zerolog.MessageFieldName = "message"
		log.Logger = log.Output(os.Stderr).Level(zerolog.InfoLevel)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Str("config", e.Name).Msg("reloading changed config file")
	})
	viper.WatchConfig()

	cmd.Execute()
}
