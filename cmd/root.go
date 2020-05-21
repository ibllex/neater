package cmd

import (
	"log"
	"neater/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "Neater",
		Short: "Quickly organize your messy files.",
		Long:  "Quickly organize your messy documents and make life better",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define flags and configure settings.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "set config file (default is neater.json in executable directory)")
}

func er(msg interface{}) {
	log.Fatal("Error: ", msg)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find executable directory
		_, dir := utils.GetExecutablePath()

		viper.AddConfigPath(dir)
		// Search config in executable directory with name "neater" (without extension).
		viper.SetConfigName("neater")
	}

	// Read in environment variables that match.
	viper.AutomaticEnv()

	// If the config file found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file: " + viper.ConfigFileUsed())
	} else {
		er(err)
	}
}
