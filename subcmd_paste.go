package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(pasteCommand)
	pasteCommand.Flags().IntP("axis-x", "X", 0, "Axis X")
	pasteCommand.Flags().IntP("axis-y", "Y", 0, "Axis Y")
	pasteCommand.Flags().IntP("axis-width", "W", 10, "Axis Width")
	pasteCommand.Flags().IntP("axis-height", "H", 10, "Axis Height")
	pasteCommand.Flags().BoolP("insert", "i", false, "Insert rectangle")
}

var pasteCommand = &cobra.Command{
	Use:   "paste",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		x, err := f.GetInt("axis-x")
		if err != nil {
			panic(err)
		}

		y, err := f.GetInt("axis-y")
		if err != nil {
			panic(err)
		}

		w, err := f.GetInt("axis-width")
		if err != nil {
			panic(err)
		}

		h, err := f.GetInt("axis-height")
		if err != nil {
			panic(err)
		}

		useInsert, err := f.GetInt("insert")
		if err != nil {
			panic(err)
		}

		fmt.Println(x, y, w, h, useInsert)
	},
}
