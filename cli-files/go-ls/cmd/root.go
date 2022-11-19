package cmd

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/cobra"
    
)

var rootCmd = &cobra.Command{
    Use:   "hugo",
    Short: "Hugo is a very fast static site generator",
    Long: `A Fast and Flexible Static Site Generator built with
                  love by spf13 and friends in Go.
                  Complete documentation is available at https://gohugo.io/documentation/`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, World!")
    },
  }

  func Execute() {
    files, e := ioutil.ReadDir(".")
    if e != nil {
       panic(e)
    }
    for _, file := range files {
       println(file.Name())
    }
  }