/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgFileName = ".toolsium"
	cfgFileType = "json"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolsium",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	cfg, err := config.LoadDefaultConfig(context.TODO())
	// 	cfg.Credentials = aws.NewCredentialsCache(&lib.MfaProvider{})

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	client := s3.NewFromConfig(cfg)
	// 	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, object := range output.Buckets {
	// 		log.Printf("Bucket=%s", aws.ToString(object.Name))
	// 	}
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%v.%v)", cfgFileName, cfgFileType))
	// Instead of profiles for now, I recommend just passing in different config files.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".toolsium" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType(cfgFileType)
		viper.SetConfigName(cfgFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
