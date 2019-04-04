package rect

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

type Relation int

const (
	RelationNone Relation = iota
	RelationPrev
	RelationNext
)

var empty = MetaRune{Value: ' ', Relation: RelationNone}

type MetaRune struct {
	Value    rune
	Relation Relation
}

type PasteConfig struct {
	X                int
	Y                int
	Padding          string
	UseInsert        bool
	IgnoreWhiteSpace bool
}

func toMetaRune(s string) (ret []MetaRune) {
	for _, c := range s {
		l := runewidth.StringWidth(string(c))
		isFullWidth := l == 2
		if isFullWidth {
			ret = append(ret, MetaRune{Value: c, Relation: RelationNext})
			ret = append(ret, MetaRune{Value: c, Relation: RelationPrev})
		} else {
			ret = append(ret, MetaRune{Value: c, Relation: RelationNone})
		}
	}
	return
}

func toMetaRunes(s []string) (ret [][]MetaRune) {
	for _, line := range s {
		ret = append(ret, toMetaRune(line))
	}
	return
}

func maxWidth(s []string) (ret int) {
	for _, v := range s {
		l := runewidth.StringWidth(v)
		if ret < l {
			ret = l
		}
	}
	return
}

func Paste(src []string, inputData []string, config PasteConfig) (ret [][]MetaRune) {
	ret = toMetaRunes(src)
	max := maxWidth(src)
	for i, line := range inputData {
		// 貼り付け開始行数をずらす
		y := config.Y
		y += i

		// もとのテキストの行数超過があった場合は空白で埋める
		var srcLine string
		if len(src) <= y {
			srcLine = strings.Repeat(" ", max)
			line = PadSpace(line, max, config)
		} else {
			srcLine = src[y]
		}
		paddedMeta := toMetaRune(line)

		// 空白の時に元のテキストを残す
		if config.IgnoreWhiteSpace {
			paddedMeta = ReplateIgnore(toMetaRune(srcLine), paddedMeta, " 　")
		}

		srcMeta := toMetaRune(srcLine)
		srcMeta = PasteLine(srcMeta, paddedMeta, config.X)
		if y < len(ret) {
			ret[y] = srcMeta
		} else {
			ret = append(ret, srcMeta)
		}
	}
	return
}

func PasteLine(src, inputData []MetaRune, x ...int) (ret []MetaRune) {
	ret = make([]MetaRune, len(src))
	copy(ret, src)

	setFunc := func(ret []MetaRune, mr MetaRune, i int) {
		switch mr.Relation {
		case RelationNone, RelationPrev:
			ret[i] = mr
		case RelationNext:
			ret[i] = MetaRune{Value: mr.Value, Relation: RelationNext}
			ret[i+1] = MetaRune{Value: mr.Value, Relation: RelationPrev}
		default:
			msg := fmt.Sprintf("illegal metarun. metarune=%v", mr)
			panic(msg)
		}
	}

	for i, mr := range inputData {
		if 0 < len(x) && i+x[0] < len(src) {
			i += x[0]
		}
		s := ret[i]
		switch s.Relation {
		case RelationNone, RelationPrev:
			setFunc(ret, mr, i)
		case RelationNext:
			ret[i+1] = MetaRune{Value: ' ', Relation: RelationNone}
			setFunc(ret, mr, i)
		default:
			msg := fmt.Sprintf("illegal relation value. relation=%v", s)
			panic(msg)
		}
	}
	return
}

func ReplateIgnore(inputData, src []MetaRune, ignore string) (ret []MetaRune) {
	if len(inputData) < 1 || len(src) < 1 {
		return inputData
	}

	ret = make([]MetaRune, len(inputData))
	copy(ret, inputData)

	for i, ir := range ret {
		if len(src) <= i {
			break
		}

		sm := src[i]
		if strings.ContainsRune(ignore, ir.Value) {
			switch sm.Relation {
			case RelationNone, RelationPrev:
				ret[i] = sm
			case RelationNext:
				if len(ret) <= i+1 {
					continue
				}
				ir2 := ret[i+1]
				if strings.ContainsRune(ignore, ir2.Value) {
					ret[i] = sm
					ret[i+1] = src[i+1]
				}
			}
		}
	}
	// NextとPrevが対になっていないものを修正
	for i := 0; i < len(ret)-1; i++ {
		if ret[i].Relation == RelationNext && ret[i+1].Relation != RelationPrev {
			ret[i] = MetaRune{Value: ' ', Relation: RelationNone}
			continue
		}
		if ret[i].Relation == RelationPrev && ret[i-1].Relation != RelationNext {
			ret[i] = MetaRune{Value: ' ', Relation: RelationNone}
			continue
		}
	}
	return
}

func PadSpace(inputData string, strWidth int, config PasteConfig) (ret string) {
	pad := config.Padding
	padIsFullWidth := runewidth.StringWidth(pad) == 2

	var leftPad string
	x := config.X
	if x%2 != 0 && padIsFullWidth {
		inputData = " " + inputData
		x--
	}
	if padIsFullWidth {
		x /= 2
	}
	leftPad += strings.Repeat(pad, x)
	inputData = leftPad + inputData

	inL := runewidth.StringWidth(inputData)
	srcL := strWidth
	if inL < srcL {
		diff := srcL - inL
		var rightPad string
		if diff%2 != 0 && padIsFullWidth {
			rightPad += " "
			diff--
		}
		if padIsFullWidth {
			diff /= 2
		}
		rightPad += strings.Repeat(pad, diff)
		inputData += rightPad
	}
	ret = inputData
	return
}
