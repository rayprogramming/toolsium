/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rayprogramming/toolsium/lib/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolsium",
	Short: "A handy set of tools for developers",
	Long: `Toolsium is desgined to allow developers to quickly access common resources they would be expected to use.
	
These tools can include default filters, easy access to start a session on an ec2 machine, and other features.
`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is [$HOME/%%USERPROFILE%%]/%v)", config.DefaultConfigDir()))
	// TODO(JR): Build a flag for AWS profile

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Debugf("cfgFile: %v", cfgFile)
	config.Configure(cfgFile)
}
