/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	
	
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cat",
	Short: "reproduces cat command using go",
	Long: ``,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) { 
	
		if len(args) == 0 {
			log.Fatal("error no arguments")
		}
		path := args[0]

		
		file, err := os.Stat(path)
		if err != nil {
			log.Fatal("error")
		}

	
		if file.IsDir() {
			log.Fatal("error is a Directory")
		}


		words, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("error")
		}

		//prints text
    	os.Stdout.Write(words)

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




