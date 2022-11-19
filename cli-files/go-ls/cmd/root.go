package cmd

import (
    "io/ioutil"
    "github.com/spf13/cobra"
    "log"
)

var rootCmd = &cobra.Command{
    Use:   "go-ls",
    Short: "list files",
    Long: `implementation of cli command ls using go`,
    Run: func(cmd *cobra.Command, args []string) {
      files, e := ioutil.ReadDir(".")
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