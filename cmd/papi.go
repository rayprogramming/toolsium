/*
Copyright © 2022 James Ray james@rayprogramming.com

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/rayprogramming/toolsium/lib"
	"github.com/spf13/cobra"
)

// papiCmd represents the papi command
var papiCmd = &cobra.Command{
	Use:   "papi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		cfg.Region = "us-east-1"
		cfg.Credentials = aws.NewCredentialsCache(&lib.MfaProvider{})

		if err != nil {
			log.Fatal(err)
		}

		client := ec2.NewFromConfig(cfg)
		params := &ec2.DescribeInstancesInput{
			Filters: []types.Filter{{
				Name:   aws.String("tag:Environment"),
				Values: []string{"beta"},
			}},
		}
		result, err := client.DescribeInstances(context.TODO(), params)

		if err != nil {
			fmt.Println("Error calling ec2: ", err)
			return
		}
		count := len(result.Reservations)
		fmt.Println("Instances: ", count)
		for i, reservation := range result.Reservations {
			for k, instance := range reservation.Instances {
				fmt.Printf("Instance number: %v - %v Id: %v \n", i, k, aws.ToString(instance.InstanceId))
			}
		}
	},
}

func init() {
	manageCmd.AddCommand(papiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// papiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// papiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
