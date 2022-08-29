/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rayprogramming/toolsium/lib/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	awsProfile  string
	skipAWSInit bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup toolsium",
	Run: func(cmd *cobra.Command, args []string) {
		if !skipAWSInit {
			runAWS()
		}
		log.Debugf("init conf path: %v", config.GetConfigFilePath())
		err := config.CreateConfig()
		cobra.CheckErr(err)
	},
}

func runAWS() {
	fmt.Println("Please configure AWS:")
	cmd := exec.Command("aws", "configure", "--profile", awsProfile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
	log.Debugln("Done running AWS command")
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&awsProfile, "awsProfile", "a", "default", "AWS Profile")
	initCmd.Flags().BoolVar(&skipAWSInit, "skip-aws", false, "Skip AWS configure")

	initCmd.Flags().StringP("mfa_serial", "m", "", "The MFA serial for your account")
	config.GetViper().BindPFlag("mfa_serial", initCmd.Flags().Lookup("mfa_serial"))
	config.GetViper().SetDefault("mfa_serial", "")

	initCmd.Flags().StringP("department", "d", "", "Department filter")
	config.GetViper().BindPFlag("department", initCmd.Flags().Lookup("department"))
	config.GetViper().SetDefault("department", "")

}
