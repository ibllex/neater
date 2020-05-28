package cmd

import (
	"log"
	"neater/parser"
	"neater/utils"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config The neater config file struct
type Config struct {
	Output    string
	Operation string
	Override  bool
	Rules     map[string][]parser.RuleConfig
	Filters   []parser.FilterConfig
}

var (
	cfgFile string
	opsPath string
	config  Config

	filterContainer = parser.FilterContainer{}

	rootCmd = &cobra.Command{
		Use:   "Neater",
		Short: "Quickly organize your messy files.",
		Long:  "Quickly organize your messy documents and make life better",
		Run:   run,
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
	viper.SetConfigType("json")
	cobra.OnInitialize(initConfig)

	// Define flags and configure settings.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "set config file (default is neater.json in executable directory)")
	rootCmd.PersistentFlags().StringVarP(&opsPath, "path", "p", ".", "set the directory to be processed")
	rootCmd.PersistentFlags().StringP("output", "o", "", "set the output path")

	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func er(msg interface{}) {
	log.Fatal("Error: ", msg)
}

// initConfig Load the config file and set global variables
func initConfig() {
	var err error
	rules := make(map[string](*parser.Rule))

	// Get the ops path
	if opsPath, err = filepath.Abs(opsPath); err != nil {
		er(err)
	}

	// Load config file
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
	if err = viper.ReadInConfig(); err != nil {
		er(err)
	}

	// Parsing configuration files
	err = viper.Unmarshal(&config)
	if err != nil {
		er(err)
	}

	// Generate all rules
	for name, rule := range config.Rules {
		rules[name] = parser.NewRule(name, rule)
	}

	// Generate all filters
	for _, config := range config.Filters {
		if rules[config.Rule] != nil {
			filterContainer.Add(parser.NewFilter(config, rules[config.Rule]))
		} else {
			log.Fatal("Rule " + config.Rule + " not found!")
		}
	}

	// Convert output path to absolute path
	if config.Output, err = filepath.Abs(config.Output); err != nil {
		er(err)
	}
}

// run print key folders and walk the ops path and process all the files in it
func run(cmd *cobra.Command, args []string) {
	log.Println("Using config file: " + viper.ConfigFileUsed())
	log.Println("The ops path is: " + opsPath)
	log.Println("The output path is: " + config.Output)

	err := filepath.Walk(opsPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if err = ops(path); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func ops(p string) error {
	if filter := filterContainer.Match(path.Ext(p)); filter != nil {
		filter.GetRule().Exec(p, config.Output, strings.ToLower(config.Operation) == "move")
	}

	return nil
}
