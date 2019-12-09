package cmd

import (
	"log"
	"os"
	"path"

	"github.com/unknwon/com"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tosone/logging"
	"github.com/tosone/mirrorepo/cmd/scan"
	"github.com/tosone/mirrorepo/cmd/version"
	"github.com/tosone/mirrorepo/cmd/web"
	"github.com/tosone/mirrorepo/common"
	"github.com/tosone/mirrorepo/models"
	"github.com/tosone/mirrorepo/services"
)

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   "mirrorepo",
	Short: "mirrorepo sync repo to remote repo",
	Long:  ``,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "search git repo in dir",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		version.Initialize()
	},
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "search git repo in directory",
	Long:  `search git repo in directory`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		initConfig()
		scan.Initialize(args...)
	},
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start web service to see all the repo information detail",
	Long:  `start web service to see all the repo information detail`,
	Run: func(_ *cobra.Command, _ []string) {
		initConfig()
		web.Initialize()
	},
}

// config command line params
var config string

func init() {
	RootCmd.PersistentFlags().StringVarP(&config, "config", "f", "./config.yml", "config file")

	RootCmd.AddCommand(scanCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(webCmd)
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(common.EnvPrefix)
	if com.IsFile(config) {
		viper.SetConfigFile(config)
	} else {
		logging.Fatal("Cannot find config file. Please check.")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logging.Panic(err)
	}

	if !com.IsDir(path.Dir(viper.GetString("Log.Filename"))) {
		if err := os.MkdirAll(path.Dir(viper.GetString("Log.Filename")), 0755); err != nil {
			log.Fatalln(err)
		}
	}

	if err := models.Connect(); err != nil {
		logging.Fatal(err)
	}

	logging.Setting(logging.Config{
		LogLevel:   logging.Level(viper.GetInt("Log.Level")),
		Filename:   viper.GetString("Log.Filename"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackups: viper.GetInt("Log.MaxBackups"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		LocalTime:  viper.GetBool("Log.LocalTime"),
		Compress:   viper.GetBool("Log.Compress"),
	})

	services.Initialize()
}
