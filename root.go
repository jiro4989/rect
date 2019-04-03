package main

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var RootCommand = &cobra.Command{
	Use:     "rect",
	Short:   "rect is text rectangle editor",
	//Example: "rect right README.md",
	Version: Version,
	Long: ``,
}
