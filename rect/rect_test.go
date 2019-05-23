package rect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasteLine(t *testing.T) {
	type TestData struct {
		desc   string
		expect string
		src    string
		instr  string
		x      int
	}
	tds := []TestData{
		{desc: "通常", expect: "abc45", src: "12345", instr: "abc", x: 0},
		{desc: "座標ズレ", expect: "1abc5", src: "12345", instr: "abc", x: 1},
		{desc: "全角", expect: "1234う", src: "あいう", instr: "1234", x: 0},
		{desc: "全角わりこみ", expect: " 123う", src: "あいう", instr: "123", x: 1},
		{desc: "全角わりこみ(末尾残り)", expect: " 1234 ", src: "あいう", instr: "1234", x: 1},
		{desc: "全角と全角", expect: "えおう", src: "あいう", instr: "えお", x: 0},
		{desc: "全角と全角のわりこみ", expect: " えお ", src: "あいう", instr: "えお", x: 1},
	}
	for _, v := range tds {
		got := PasteLine(v.src, v.instr, v.x)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
