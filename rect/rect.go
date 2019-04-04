package rect

import "github.com/mattn/go-runewidth"

type PasteConfig struct {
	X                int
	Y                int
	Padding          string
	UseInsert        bool
	IgnoreWhiteSpace bool
}

func Paste(src, inputData []string, config PasteConfig) (ret []string) {
	return
}

func PasteLine(src, inputData string, config PasteConfig) string {
	// マルチバイト文字を扱うためにruneに変換
	srcRune := []rune(src)
	inputRune := []rune(inputData)
	// srcを直接上書きするためにrangeを使わない
	for i := 0; i < len(inputRune); i++ {
		ir := inputRune[i]
		// 座標ズレ分を加算
		// 直接i2を加算するとループ回数が変わってしまうため一時変数を使用
		i2 := i + config.X

		// 上書き元のテキストの範囲超過していたら空文字を追加
		if len(srcRune) <= i2 {
			srcRune = append(srcRune, ' ')
		}
		sr := srcRune[i2]
		isFullWidth := runewidth.StringWidth(string(sr)) == 2
		if isFullWidth {
			// 変更対象が全角かつ、上書きに使う値が全角のときはそのまま置き換える
			srcRune[i2] = ir

			// 変更対象が全角かつ、上書きに使う値が半角のときは
			if runewidth.StringWidth(string(ir)) == 2 {
				continue
			}
			nextIndex := i2 + 1
			if nextIndex < len(srcRune) {
				srcRune[i2+1] = ir
			}
			continue
		}

		// 変更対象が半角のとき、上書きに使う値が全角だった場合は
		// その次の値も更新する
		srcRune[i2] = ir
		if runewidth.StringWidth(string(ir)) == 2 {
			nextIndex := i2 + 1
			if nextIndex < len(srcRune) {
				srcRune = append(srcRune[:nextIndex], srcRune[nextIndex+1:]...)
			}
		}
	}
	return string(srcRune)
}
