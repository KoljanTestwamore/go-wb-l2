/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

const (
	FIELDS = "FIELDS"
	DELIMETER = "delimeter"
	ONLY_DELIMETED = "only-delimeted"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cut",
	Short: "remove sections from each line of files",
	Long: `Print selected parts of lines from STDIO to standard output.

       Mandatory arguments to long options are mandatory for short
       options too.`,
	Run: func(cmd *cobra.Command, args []string) { 
		f, _ := cmd.Flags().GetUintSlice(FIELDS)
		d, _ := cmd.Flags().GetString(DELIMETER)
		s, _ := cmd.Flags().GetBool(ONLY_DELIMETED)

		reader := bufio.NewReader(os.Stdin)

		for {
			text, _ := reader.ReadString('\n')
			cut := strings.Split(text, d)
			res := ""

			if s && len(cut) == 1 {
				continue
			}

			for i, word := range cut {
				if (slices.Contains(f, uint(i + 1))) {
					res += word + " "
				}
			}

			fmt.Println(res)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cut.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().UintSliceP(FIELDS, "f", []uint{}, "select only these fields;")
	rootCmd.Flags().StringP(DELIMETER, "d", "\t", "select only these fields;")
	rootCmd.Flags().BoolP(ONLY_DELIMETED, "s", false, "select only these fields;")
}


