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
	width := config.X
	for i := 0; i < len(inputRune); i++ {
		// 上書き元のテキストの表示幅を超過シていたら追加
		if runewidth.StringWidth(src) <= width {
			src += " "
			srcRune = append(srcRune, ' ')
		}

		ir := inputRune[i]
		w := runewidth.RuneWidth(ir)

		// 座標ズレ分を加算
		// 表示上の位置から処理対象の文字とそのインデックスを取得
		sr, i2 := lookup(src, width)

		isFullWidth := runewidth.RuneWidth(sr) == 2
		if isFullWidth {
			// 変更対象が全角かつ、上書きに使う値が全角のときはそのまま置き換える
			srcRune[i2] = ir
			// 全角置き換えのときに半角１文字分ずれて置換したなら
			// 置き換えた文字の直前に半角スペースを追加する
			// _, i3 := lookup(src, width-1)
			// if i2 == i3 {
			// 	a := srcRune[:i2-1]
			// 	b := srcRune[i2-1:]
			// 	srcRune = append(a, ' ')
			// 	srcRune = append(srcRune, b...)
			// }
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
