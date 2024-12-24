/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task2.9",
	Short: "Downloads website",
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) { 
		url := args[0]
		pwd, err := os.Getwd()
		
		if err != nil {
			panic(err.Error())
		}

		resp, err := http.Get("https://" + url)
		if err != nil {
			panic(err.Error())
		}

		fo, err := os.Create(pwd + "/index.html")
		if err != nil {
			panic(err.Error())
		}
		defer fo.Close()

		body, _ := io.ReadAll(resp.Body)
		_, err = fo.Write(body)
		if err != nil {
			panic(err.Error())
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task2.9.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


