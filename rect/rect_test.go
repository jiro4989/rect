package rect

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLookup(t *testing.T) {
	type TestData struct {
		desc, s string
		expect  rune
		expect2 int
		n       int
	}
	testdatas := []TestData{
		{desc: "半角のみ+先頭の値", expect: '1', expect2: 0, s: "12345", n: 0},
		{desc: "半角のみ+途中の値", expect: '2', expect2: 1, s: "12345", n: 1},
		{desc: "半角のみ+末尾の値", expect: '5', expect2: 4, s: "12345", n: 4},
		{desc: "半角のみ負の値", expect: rune(0), expect2: 0, s: "12345", n: -1},
		{desc: "半角のみ範囲外", expect: rune(0), expect2: 0, s: "12345", n: 5},
		{desc: "全角のみ+0", expect: 'あ', expect2: 0, s: "あいうえお", n: 0},
		{desc: "全角のみ+1", expect: 'あ', expect2: 0, s: "あいうえお", n: 1},
		{desc: "全角のみ+2(次の文字)", expect: 'い', expect2: 1, s: "あいうえお", n: 2},
		{desc: "全角のみ最後の文字", expect: 'お', expect2: 4, s: "あいうえお", n: 8},
		{desc: "全角のみ最後の文字", expect: 'お', expect2: 4, s: "あいうえお", n: 9},
		{desc: "全角のみ範囲外", expect: rune(0), expect2: 0, s: "あいうえお", n: 10},
		{desc: "全角半角の混在", expect: 'a', expect2: 0, s: "aあ", n: 0},
		{desc: "全角半角の混在", expect: 'あ', expect2: 1, s: "aあ", n: 1},
		{desc: "全角半角の混在", expect: 'あ', expect2: 1, s: "aあ", n: 2},
		{desc: "全角半角の混在(範囲外)", expect: rune(0), expect2: 0, s: "aあ", n: 3},
	}
	for _, v := range testdatas {
		got, index := lookup(v.s, v.n)
		if diff := cmp.Diff(v.expect, got); diff != "" {
			msg := fmt.Sprintf("NG %s\n%s", v.desc, diff)
			t.Error(msg)
		} else {
			msg := fmt.Sprintf("OK %s", v.desc)
			t.Log(msg)
		}
		if diff := cmp.Diff(v.expect2, index); diff != "" {
			msg := fmt.Sprintf("NG %s\n%s", v.desc, diff)
			t.Error(msg)
		} else {
			msg := fmt.Sprintf("OK %s", v.desc)
			t.Log(msg)
		}
	}
}

func TestPaste(t *testing.T) {
	type TestData struct {
		desc                   string
		expect, src, inputData []string
		config                 PasteConfig
	}
	testdatas := []TestData{
		{desc: "半角のみ1行", expect: []string{"abc45"}, src: []string{"12345"}, inputData: []string{"abc"}, config: PasteConfig{}},
		{desc: "半角のみ2行、１行更新１行更新しない", expect: []string{"abc45", "67890"}, src: []string{"12345", "67890"}, inputData: []string{"abc"}, config: PasteConfig{}},
		{desc: "半角のみ2行", expect: []string{"abc45", "def90"}, src: []string{"12345", "67890"}, inputData: []string{"abc", "def"}, config: PasteConfig{}},
		{desc: "半角のみ2行+X:1", expect: []string{"1abc5", "6def0"}, src: []string{"12345", "67890"}, inputData: []string{"abc", "def"}, config: PasteConfig{X: 1}},
		{desc: "半角のみ2行+X:1,Y:1", expect: []string{"12345", "6abc0", " def "}, src: []string{"12345", "67890"}, inputData: []string{"abc", "def"}, config: PasteConfig{X: 1, Y: 1}},
		{desc: "全角のみ2行", expect: []string{"さしうえお", "すせそけこ"}, src: []string{"あいうえお", "かきくけこ"}, inputData: []string{"さし", "すせそ"}, config: PasteConfig{}},
		{desc: "全角のみ2行+X:1", expect: []string{"さしうえお", "すせそけこ"}, src: []string{"あいうえお", "かきくけこ"}, inputData: []string{"さし", "すせそ"}, config: PasteConfig{X: 1}},
		{desc: "全角のみ2行+X:2", expect: []string{"あさしえお", "かすせそこ"}, src: []string{"あいうえお", "かきくけこ"}, inputData: []string{"さし", "すせそ"}, config: PasteConfig{X: 2}},
		//{desc: "全角のみ2行+X:2,Y:1", expect: []string{"あいうえお", "かさしそこ", " すせそ   "}, src: []string{"あいうえお", "かきくけこ"}, inputData: []string{"さし", "すせそ"}, config: PasteConfig{X: 2, Y: 1}},
	}
	for _, v := range testdatas {
		got := Paste(v.src, v.inputData, v.config)
		if diff := cmp.Diff(v.expect, got); diff != "" {
			msg := fmt.Sprintf("NG %s\n%s", v.desc, diff)
			t.Error(msg)
		} else {
			msg := fmt.Sprintf("OK %s", v.desc)
			t.Log(msg)
		}
	}
}

func TestPasteLine(t *testing.T) {
	type TestData struct {
		desc, expect, src, inputData string
		config                       PasteConfig
	}
	testdatas := []TestData{
		{desc: "半角のみ原点座標開始", expect: "abc45", src: "12345", inputData: "abc", config: PasteConfig{}},
		{desc: "半角のみ原点+1X座標開始", expect: "1abc5", src: "12345", inputData: "abc", config: PasteConfig{X: 1}},
		{desc: "半角のみ原点ズレによる元テキストの範囲超過", expect: "123abc", src: "12345", inputData: "abc", config: PasteConfig{X: 3}},
		{desc: "全角あり原点座標開始", expect: "あい5", src: "12345", inputData: "あい", config: PasteConfig{}},
		{desc: "全角あり原点+1X座標開始", expect: "1あい", src: "12345", inputData: "あい", config: PasteConfig{X: 1}},
		{desc: "全角が全角を置き換える", expect: "あいお", src: "うえお", inputData: "あい", config: PasteConfig{}},
		{desc: "全角が元テキストの範囲超過", expect: "あい", src: "お", inputData: "あい", config: PasteConfig{}},
		{desc: "全角文字を半角文字で置き換える", expect: "a２３", src: "１２３", inputData: "a", config: PasteConfig{}},
		// {desc: "全角文字の途中を始点に、全角文字を半角文字で置き換える", expect: " a２３", src: "１２３", inputData: "a", config: PasteConfig{X: 1}},
		// {desc: "全角文字の途中を始点に、全角文字を半角文字で置き換える", expect: "１ a３", src: "１２３", inputData: "a", config: PasteConfig{X: 3}},
		//{desc: "全角は半角２文字分で全角１文字分ずれる", expect: "おあい", src: "お", inputData: "あい", config: PasteConfig{X: 2}},
	}
	for _, v := range testdatas {
		got := PasteLine(v.src, v.inputData, v.config)
		if diff := cmp.Diff(v.expect, got); diff != "" {
			msg := fmt.Sprintf("NG %s\n%s", v.desc, diff)
			t.Error(msg)
		} else {
			msg := fmt.Sprintf("OK %s", v.desc)
			t.Log(msg)
		}
	}
}
