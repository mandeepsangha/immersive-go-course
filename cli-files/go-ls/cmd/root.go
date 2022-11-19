package cmd

import (
    "github.com/spf13/cobra"
    "log"
    "os"
    "fmt"
  
)

var rootCmd = &cobra.Command{
    Use:   "go-ls",
    Short: "list files",
    Long: ` cli command ls using go`,
    Args: cobra.ArbitraryArgs,
    Run: func(cmd *cobra.Command, args []string) {
     var directory =""
      if len(args) == 0 {
        // if no argument set directory to current directory
        directory= "."
      } else {
        // if argument set directory to argument
        directory = args[0]
      }

      fileInfo, err := os.Stat(directory)
        if err != nil {
          log.Fatal(err)
        }
        if !fileInfo.IsDir() {
          fmt.Printf("%s\n", directory)
          return
        }

      // Read Directory file names
      files, err := os.ReadDir(directory)
      if err != nil {
         log.Fatal(err)
      }
      for _, file := range files {
        // for loop to print each file name
         println(file.Name())
      }
    },
  }

  func Execute() {
    if err := rootCmd.Execute(); err != nil {
      log.Fatal(err)
    }
  }