package rect

import "strings"

type PasteConfig struct {
	X                int
	Y                int
	Padding          string
	UseInsert        bool
	IgnoreWhiteSpace bool
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

func Paste(editTarget []string, pasteData []string, config PasteConfig) (ret []string) {
	runes := toRunes(editTarget)
	emptyLine := []rune(strings.Repeat(" ", len(runes[0])))
	for y, line := range pasteData {
		y += config.Y
		if len(runes) <= y {
			runes = append(runes, emptyLine)
		}
		for x, c := range line {
			x += config.X
			runes[y][x] = c
		}
	}
	for _, r := range runes {
		s := string(r)
		ret = append(ret, s)
	}
	return
}
