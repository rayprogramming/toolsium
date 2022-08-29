/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"github.com/rayprogramming/toolsium/lib/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup toolsium",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("init conf path: %v", config.GetConfigFilePath())
		err := config.CreateConfig()
		cobra.CheckErr(err)
	},
}

func runAWS() {
	log.Debugln("Done running AWS command")
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("mfa_serial", "m", "", "The MFA serial for your account")
	config.GetViper().BindPFlag("mfa_serial", initCmd.Flags().Lookup("mfa_serial"))
	config.GetViper().SetDefault("mfa_serial", "")

	initCmd.Flags().StringP("department", "d", "", "Department filter")
	config.GetViper().BindPFlag("department", initCmd.Flags().Lookup("department"))
	config.GetViper().SetDefault("department", "")
}
