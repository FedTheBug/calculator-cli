/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetBool("float")
		if status {
			result, err := multiplyFloat(args)
			if err != nil {
				fmt.Printf("Error in calculating Float multiplication %v", err)
			} else {
				fmt.Printf("Multiplication of %s is %.3f", args, result)
			}
		} else {
			result, err := multiplyInt(args)
			if err != nil {
				fmt.Printf("Error in calculating Int multiplication %v", err)
			} else {
				fmt.Printf("Multiplication of %s is %d", args, result)
			}
		}
	},
}

func multiplyInt(s []string) (int, error) {
	product := 1
	for _, i := range s {
		val, err := strconv.Atoi(i)
		if err != nil {
			return product, err
		}
		product = product * val
	}
	return product, nil
}

func multiplyFloat(s []string) (float64, error) {
	product := 1.0
	for _, i := range s {
		val, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return product, err
		}
		product = product * val
	}
	return product, nil
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
	multiplyCmd.Flags().BoolP("float", "f", false, "Multiplication of Floating Point Numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiplyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
