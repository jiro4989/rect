package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jiro4989/rect/rect"
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
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "need over 1 argument.")
			fmt.Fprintln(os.Stderr, "see `rect -h`")
			os.Exit(1)
		}

		f := cmd.Flags()

		x, err := f.GetInt("axis-x")
		if err != nil {
			panic(err)
		}

		y, err := f.GetInt("axis-y")
		if err != nil {
			panic(err)
		}

		// w, err := f.GetInt("axis-width")
		// if err != nil {
		// 	panic(err)
		// }
		//
		// h, err := f.GetInt("axis-height")
		// if err != nil {
		// 	panic(err)
		// }
		//
		// useInsert, err := f.GetInt("insert")
		// if err != nil {
		// 	panic(err)
		// }

		config := rect.PasteConfig{X: x, Y: y}

		// 入力データ取得
		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		inputDatas := strings.Split(string(b), "\n")

		// 加工対象のテキスト取得
		var src []string
		if len(args) < 2 {
			src = readStdin()
		} else {
			b, err := ioutil.ReadFile(args[1])
			if err != nil {
				panic(err)
			}
			src = strings.Split(string(b), "\n")
		}

		ret := rect.Paste(src, inputDatas, config)
		for _, line := range ret {
			fmt.Println(line)
		}
	},
}
