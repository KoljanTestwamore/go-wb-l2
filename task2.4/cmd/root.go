/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	KEY = "toggle"
	REVERSE = "reverse"
	UNIQUE = "unique"
	NUMERIC_SORT = "numeric-sort"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task24",
	Short: "sort - sort lines of text files",
	Long: `Write sorted concatenation of all FILE to standard output.

	Mandatory arguments to long options are mandatory for short
	options too.  Ordering options:`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		K, _ := cmd.Flags().GetUint(KEY)
		R, _ := cmd.Flags().GetBool(REVERSE)
		U, _ := cmd.Flags().GetBool(UNIQUE)
		NS, _ := cmd.Flags().GetBool(NUMERIC_SORT)

		path := args[0]

		uniques := make(map[string]bool)

		pwd, _ := os.Getwd()

		file, err := os.Open(pwd + "/" + path)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		fileStrings := []string{}

		for scanner.Scan() {
			text := scanner.Text()

			if U {
				element := uniques[text]
				if element {
					continue
				}
				uniques[text] = true
			}

			fileStrings = append(fileStrings, text)
		}

		sort.SliceStable(fileStrings, func (i, j int) bool {
			str1 := fileStrings[i]
			str2 := fileStrings[j]

			if K != 0 {
				str1 = strings.Fields(str1)[K]
				str2 = strings.Fields(str2)[K]
			}

			if NS {
				val1, err1 := strconv.Atoi(str1)
				val2, err2 := strconv.Atoi(str2)
				if err1 != nil || err2 != nil {
					return false
				}

				return val1 > val2
			
			}

			return strings.Compare(str1, str2) == 1
		})

		if R {
			for _, str := range fileStrings {
				fmt.Println(str)
			}
			return
		}

		for i := range fileStrings {
			fmt.Println(fileStrings[len(fileStrings) - i - 1])
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task24.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().UintP(KEY, "k", 0, "sort via a key; KEYDEF gives location and type")
	rootCmd.Flags().BoolP(REVERSE, "r", false, "reverse the result of comparisons")
	rootCmd.Flags().BoolP(UNIQUE, "u", false, "with -c, check for strict ordering; without -c, output only the first of an equal run")
	rootCmd.Flags().BoolP(NUMERIC_SORT, "n", false, "compare according to string numerical value; see manual for which strings are supported")
}


