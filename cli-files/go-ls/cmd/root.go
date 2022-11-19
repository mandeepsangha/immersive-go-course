package cmd

import (
    "io/ioutil"
    "github.com/spf13/cobra"
    "log"
)

var rootCmd = &cobra.Command{
    Use:   "go-ls",
    Short: "list files",
    Long: ` cli command ls using go`,
    Args: cobra.ArbitraryArgs,
    Run: func(cmd *cobra.Command, args []string) {
     var directory =""
      if len(args) == 0 {
        directory= "."
      } else {
        directory = args[0]
      }
      files, e := ioutil.ReadDir(directory)
      if e != nil {
         panic(e)
      }
      for _, file := range files {
         println(file.Name())
      }
    },
  }

  func Execute() {
    if err := rootCmd.Execute(); err != nil {
      log.Fatal(err)
    }
  }