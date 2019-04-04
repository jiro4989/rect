package rect

import "github.com/mattn/go-runewidth"

type PasteConfig struct {
	X                int
	Y                int
	Padding          string
	UseInsert        bool
	IgnoreWhiteSpace bool
}

func lookup(s string, n int) (ret rune, index int) {
	if n < 0 || len(s) < 1 || runewidth.StringWidth(s) <= n {
		return
	}
	var width int
	for i, v := range []rune(s) {
		w := runewidth.RuneWidth(v)
		width += w
		if n < width {
			ret = v
			index = i
			return
		}
	}
	return
}

func Paste(src, inputData []string, config PasteConfig) (ret []string) {
	return
}

func PasteLine(src, inputData string, config PasteConfig) string {
	// マルチバイト文字を扱うためにruneに変換
	srcRune := []rune(src)
	inputRune := []rune(inputData)
	// srcを直接上書きするためにrangeを使わない
	var width int
	for i := 0; i < len(inputRune); i++ {
		ir := inputRune[i]
		// 座標ズレ分を加算
		// 直接i2を加算するとループ回数が変わってしまうため一時変数を使用
		w := runewidth.RuneWidth(ir)

		// 表示上の位置から処理対象の文字とそのインデックスを取得
		sr, i2 := lookup(src, width)
		i2 += config.X

		// 上書き元のテキストの範囲超過していたら空文字を追加
		if len(srcRune) <= i2 {
			srcRune = append(srcRune, ' ')
		}
		isFullWidth := runewidth.RuneWidth(sr) == 2
		if isFullWidth {
			// 変更対象が全角かつ、上書きに使う値が全角のときはそのまま置き換える
			srcRune[i2] = ir
			width += w
			continue
		}

		// 変更対象が半角のとき、上書きに使う値が全角だった場合は
		// その次の値も更新する
		srcRune[i2] = ir
		if runewidth.RuneWidth(ir) == 2 {
			nextIndex := i2 + 1
			if nextIndex < len(srcRune) {
				srcRune = append(srcRune[:nextIndex], srcRune[nextIndex+1:]...)
				width--
			}
		}
		width += w
	}
	return string(srcRune)
}
