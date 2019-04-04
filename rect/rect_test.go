package rect

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
