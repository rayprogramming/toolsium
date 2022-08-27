/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mfaCmd represents the mfa command
var mfaCmd = &cobra.Command{
	Use:   "mfa [token]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires mfa token")
		}
		if len(args[0]) < 6 || len(args[0]) > 6 {
			return errors.New("invalid token")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		client := sts.NewFromConfig(cfg)
		sessionTokenInput := sts.GetSessionTokenInput{
			SerialNumber: aws.String(viper.GetString("mfa_serial")),
			TokenCode:    aws.String(args[0]),
		}
		tokenOutput, err := client.GetSessionToken(context.TODO(), &sessionTokenInput)
		if err != nil {
			log.Fatal(err)
		}
		viper.Set("session", tokenOutput)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mfaCmd)
}
