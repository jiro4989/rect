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

func toRunes(s []string) (ret [][]rune) {
	for _, line := range s {
		var buf []rune
		for _, c := range line {
			buf = append(buf, c)
		}
		ret = append(ret, buf)
	}
	return
}

func Paste(src []string, inputData []string, config PasteConfig) (ret []rune) {
	// pad := config.Padding
	// X := config.X
	// padLen := runewidth.StringWidth(pad)
	// useFullWidthPadding := padLen < 2
	//
	// srcMetaRunes := toMetaRune(src)
	// inputMetaRunes := toMetaRune(inputData)
	//
	// for i, mr := range inputMetaRunes {
	// 	if i < X {
	// 		continue
	// 	}
	// 	srcM := srcMetaRunes[i]
	// }
	return
}

func PasteLine(src []MetaRune, inputData []MetaRune) (ret []MetaRune) {
	ret = make([]MetaRune, len(src))
	copy(ret, src)

	setFunc := func(ret []MetaRune, mr MetaRune, i int) {
		switch mr.Relation {
		case RelationNone:
			ret[i] = mr
		case RelationPrev:
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
		s := ret[i]
		switch s.Relation {
		case RelationNone:
			setFunc(ret, mr, i)
		case RelationPrev:
			// ret[i-1] = MetaRune{Value: ' ', Relation: RelationNone}
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

func PadSpace(src, inputData string, config PasteConfig) (ret string) {
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
	srcL := runewidth.StringWidth(src)
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

// func Paste(editTarget []string, pasteData []string, config PasteConfig) (ret []string) {
// 	runes := toRunes(editTarget)
// 	emptyLine := []rune(strings.Repeat(" ", len(runes[0])))
// 	for y, line := range pasteData {
// 		y += config.Y
// 		if len(runes) <= y {
// 			runes = append(runes, emptyLine)
// 		}
// 		for x, c := range line {
// 			x += config.X
// 			runes[y][x] = c
// 		}
// 	}
// 	for _, r := range runes {
// 		s := string(r)
// 		ret = append(ret, s)
// 	}
// 	return
// }
