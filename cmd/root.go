package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/mirror-repo/cmd/scan"
	"github.com/tosone/mirror-repo/cmd/version"
	"github.com/tosone/mirror-repo/cmd/web"
	"github.com/tosone/mirror-repo/logging"
)

var cfgFile string

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   "Mirror-repo",
	Short: "Mirror-repo sync repo to remote repo",
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
		scan.Initialize(args)
	},
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start web service to see all the repo information detail",
	Long:  `start web service to see all the repo information detail`,
	Run: func(_ *cobra.Command, _ []string) {
		web.Initialize()
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "/etc/mirror-repo/config.yaml", "config file")

	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(scanCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(webCmd)
}

func initConfig() {
	defaultConfig()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("/etc/mirror-repo/config.yaml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Cannot find the special config file.")
	}
	logging.Setting()
}

func defaultConfig() {
	viper.SetDefault("DatabaseEngine", "sqlite3")
	viper.SetDefault("log", "err.log")
}
