package rect

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

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
	ret = make([]string, len(src))
	copy(ret, src)

	for i, v := range inputData {
		y := i + config.Y
		if len(ret) <= y {
			w := runewidth.StringWidth(src[0])
			s := strings.Repeat(" ", w)
			ret = append(ret, s)
		}
		ret[y] = PasteLine(ret[y], v, config)
	}
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
		srcWidth := runewidth.StringWidth(src)
		if srcWidth <= width {
			pad := strings.Repeat(" ", width-srcWidth+1)
			src += pad
			srcRune = append(srcRune, []rune(pad)...)
		}

		ir := inputRune[i]
		w := runewidth.RuneWidth(ir)

		// 表示上の位置から処理対象の文字とそのインデックスを取得
		sr, i2 := lookup(src, width)

		if len(srcRune) <= i2 {
			n := 1
			if runewidth.RuneWidth(sr) == 2 {
				n = 2
			}
			pad := strings.Repeat(" ", n)
			src += pad
			srcRune = append(srcRune, []rune(pad)...)
		}

		isFullWidth := runewidth.RuneWidth(sr) == 2
		if isFullWidth {
			// 変更対象が全角かつ、上書きに使う値が全角のときはそのまま置き換える
			srcRune[i2] = ir
			// 全角置き換えのときに半角１文字分ずれて置換したなら
			// 置き換えた文字の直前に半角スペースを追加する
			// _, i3 := lookup(src, width-1)
			// if i2 == i3 {
			// 	v := append([]rune{' '}, srcRune[i2-1:]...)
			// 	srcRune = append(srcRune[:i2-1], v...)
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
