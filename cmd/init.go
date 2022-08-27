/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup toolsium",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		if err := viper.WriteConfigAs(filepath.Join(home, cfgFileName)); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("mfa_serial", "m", "", "The MFA serial for your account")
	viper.BindPFlag("mfa_serial", initCmd.Flags().Lookup("mfa_serial"))
	viper.SetDefault("mfa_serial", "")

	initCmd.Flags().StringP("department", "d", "", "Department filter")
	viper.BindPFlag("department", initCmd.Flags().Lookup("department"))
	viper.SetDefault("department", "")
}
