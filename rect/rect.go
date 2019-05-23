package rect

import (
	"github.com/mattn/go-runewidth"
)

func PasteLine(src, instr string, x int) string {
	var (
		runeSrc    = []rune(src)
		instrWidth = runewidth.StringWidth(instr)
		instrIndex int
		pos        int
		ret        []rune
	)
	for i := 0; i < len(runeSrc); i++ {
		v := runeSrc[i]
		w := runewidth.RuneWidth(v)
		if x <= pos && pos < instrWidth+x {
			for j := 0; j < w; j++ {
				r := []rune(instr)[instrIndex]
				ret = append(ret, r)
				instrIndex++
			}
			pos += w
			continue
		}
		ret = append(ret, v)
		pos += w
	}
	return string(ret)
}
