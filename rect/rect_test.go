package rect

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPasteLine(t *testing.T) {
	type TestData struct {
		expect, src, inputData string
		config                 PasteConfig
	}
	testdatas := []TestData{
		{expect: "abc45", src: "12345", inputData: "abc", config: PasteConfig{}},
		{expect: "1abc5", src: "12345", inputData: "abc", config: PasteConfig{X: 1}},
		{expect: "123abc", src: "12345", inputData: "abc", config: PasteConfig{X: 3}},
	}
	for _, v := range testdatas {
		got := PasteLine(v.src, v.inputData, v.config)
		if diff := cmp.Diff(v.expect, got); diff != "" {
			t.Error("NG\n" + diff)
		} else {
			t.Log("OK")
		}
	}
}
