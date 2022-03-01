/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "askcloudhelloworld",
	Short: "A basic Hello World example",
	Long: `
	A longer description on how this works
	example: askcloudhelloworld
	example: askcloudhelloworld -d
	`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetBool("differentmessage")
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("Hello World Alternative!")
			return
		}
		fmt.Println("Hello World!")
	},
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
	rootCmd.Flags().BoolP("differentmessage", "d", false, "Toggle a different message")
}
