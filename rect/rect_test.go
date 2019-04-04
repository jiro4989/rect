package rect

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestToMetaRunes(t *testing.T) {
	assert.Equal(t, [][]MetaRune{
		{
			{Value: 'あ', Relation: RelationNext},
			{Value: 'あ', Relation: RelationPrev},
			{Value: 'い', Relation: RelationNext},
			{Value: 'い', Relation: RelationPrev},
		},
		{
			{Value: 'う', Relation: RelationNext},
			{Value: 'う', Relation: RelationPrev},
			{Value: 'え', Relation: RelationNext},
			{Value: 'え', Relation: RelationPrev},
		},
		{
			{Value: ' ', Relation: RelationNone},
			{Value: 'あ', Relation: RelationNext},
			{Value: 'あ', Relation: RelationPrev},
			{Value: ' ', Relation: RelationNone},
		},
	}, toMetaRunes([]string{"あい", "うえ", " あ "}), "MultiByte")
	assert.Equal(t, [][]MetaRune{
		{
			{Value: 'a', Relation: RelationNone},
			{Value: 'b', Relation: RelationNone},
		}}, toMetaRunes([]string{"ab"}), "SingleByte")
}

func TestPasteLine(t *testing.T) {
	assert.Equal(t, []MetaRune{
		{Value: 'a', Relation: RelationNone},
		{Value: 'b', Relation: RelationNone},
		{Value: 'c', Relation: RelationNone},
		{Value: ' ', Relation: RelationNone},
	},
		PasteLine(toMetaRune("あい"), toMetaRune("abc")))

	assert.Equal(t, []MetaRune{
		{Value: ' ', Relation: RelationNone},
		{Value: 'a', Relation: RelationNone},
		{Value: 'b', Relation: RelationNone},
		{Value: 'c', Relation: RelationNone},
	},
		PasteLine(toMetaRune("あい"), toMetaRune(" abc")))

	assert.Equal(t, []MetaRune{
		{Value: ' ', Relation: RelationNone},
		{Value: 'あ', Relation: RelationNext},
		{Value: 'あ', Relation: RelationPrev},
		{Value: ' ', Relation: RelationNone},
	},
		PasteLine(toMetaRune("あい"), toMetaRune(" あ ")))

	assert.Equal(t, []MetaRune{
		{Value: ' ', Relation: RelationNone},
		{Value: 'あ', Relation: RelationNext},
		{Value: 'あ', Relation: RelationPrev},
		{Value: ' ', Relation: RelationNone},
	},
		PasteLine(toMetaRune("あい"), toMetaRune(" あ ")))

	assert.Equal(t, []MetaRune{
		{Value: '　', Relation: RelationNext},
		{Value: '　', Relation: RelationPrev},
		{Value: 'a', Relation: RelationNone},
		{Value: 'あ', Relation: RelationNext},
		{Value: 'あ', Relation: RelationPrev},
	},
		PasteLine(toMetaRune("12345"), toMetaRune("　aあ")))

	assert.Equal(t, []MetaRune{
		{Value: '1', Relation: RelationNone},
		{Value: 'a', Relation: RelationNone},
		{Value: 'b', Relation: RelationNone},
		{Value: 'c', Relation: RelationNone},
		{Value: '5', Relation: RelationNone},
	},
		PasteLine(toMetaRune("12345"), toMetaRune("abc"), 1))

}

func TestReplaceIgnore(t *testing.T) {
	assert.Equal(t, toMetaRune("ab1de"), ReplateIgnore(toMetaRune("  1  "), toMetaRune("abcde"), " 　"))
	assert.Equal(t, toMetaRune("ab1de"), ReplateIgnore(toMetaRune("　1　"), toMetaRune("abcde"), " 　"))
	assert.Equal(t, toMetaRune("あ1い"), ReplateIgnore(toMetaRune("  1  "), toMetaRune("あcい"), " 　"))
	assert.Equal(t, toMetaRune("あ1い"), ReplateIgnore(toMetaRune("　1　"), toMetaRune("あcい"), " 　"))
	assert.Equal(t, toMetaRune("1 123"), ReplateIgnore(toMetaRune("　1　"), toMetaRune("1あ23"), " 　"))
	assert.Equal(t, toMetaRune("1 1い"), ReplateIgnore(toMetaRune("　1　"), toMetaRune("1あい"), " 　"))
	assert.Equal(t, toMetaRune("121 3"), ReplateIgnore(toMetaRune("　1　"), toMetaRune("12あ3"), " 　"))
	assert.Equal(t, toMetaRune("abc1efg"), ReplateIgnore(toMetaRune(" 　1　 "), toMetaRune("abcdefg"), " 　"))
	assert.Equal(t, toMetaRune("abcあefg"), ReplateIgnore(toMetaRune(" 　あ　 "), toMetaRune("abcddefg"), " 　"))
	assert.Equal(t, toMetaRune(" 1"), ReplateIgnore(toMetaRune(" 1"), toMetaRune("あ"), " 　"))
	assert.Equal(t, toMetaRune(" "), ReplateIgnore(toMetaRune(" "), toMetaRune("あ"), " 　"))
	assert.Equal(t, toMetaRune("あ"), ReplateIgnore(toMetaRune("  "), toMetaRune("あ"), " 　"))
	assert.Equal(t, toMetaRune("あ"), ReplateIgnore(toMetaRune("　"), toMetaRune("あ"), " 　"))
	assert.Equal(t, toMetaRune(""), ReplateIgnore(toMetaRune(""), toMetaRune("あ"), " 　"))
	assert.Equal(t, toMetaRune("  "), ReplateIgnore(toMetaRune("  "), toMetaRune(""), " 　"))
}

func TestPadSpace(t *testing.T) {
	assert.Equal(t, "abc  ", PadSpace("abc", 5, PasteConfig{X: 0, Padding: " "}))
	assert.Equal(t, " abc ", PadSpace("abc", 5, PasteConfig{X: 1, Padding: " "}))
	assert.Equal(t, "  abc", PadSpace("abc", 5, PasteConfig{X: 2, Padding: " "}))
	assert.Equal(t, "   abc", PadSpace("abc", 5, PasteConfig{X: 3, Padding: " "}))

	assert.Equal(t, "abc 　", PadSpace("abc", 6, PasteConfig{X: 0, Padding: "　"}))
	assert.Equal(t, "abc　", PadSpace("abc", 5, PasteConfig{X: 0, Padding: "　"}))
	assert.Equal(t, " abc ", PadSpace("abc", 5, PasteConfig{X: 1, Padding: "　"}))
	assert.Equal(t, "　abc", PadSpace("abc", 5, PasteConfig{X: 2, Padding: "　"}))
	assert.Equal(t, "　 abc", PadSpace("abc", 5, PasteConfig{X: 3, Padding: "　"}))
}

func TestPaste(t *testing.T) {
	type TestData struct {
		expect         [][]MetaRune
		src, inputData []string
		config         PasteConfig
		msg            string
	}
	testDatas := []TestData{
		{
			expect: [][]MetaRune{
				{
					{Value: 'a', Relation: RelationNone},
					{Value: 'b', Relation: RelationNone},
					{Value: 'c', Relation: RelationNone},
					{Value: '4', Relation: RelationNone},
					{Value: '5', Relation: RelationNone},
				},
				{
					{Value: 'd', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
					{Value: 'f', Relation: RelationNone},
					{Value: '9', Relation: RelationNone},
					{Value: '0', Relation: RelationNone},
				},
				{
					{Value: 'g', Relation: RelationNone},
					{Value: 'h', Relation: RelationNone},
					{Value: 'i', Relation: RelationNone},
					{Value: 'd', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
				},
			},
			src:       []string{"12345", "67890", "abcde"},
			inputData: []string{"abc", "def", "ghi"},
			config:    PasteConfig{Padding: " "},
			msg:       "原点座標からの貼付け（矩形は貼り付け元の行数以内）",
		},
		{
			expect: [][]MetaRune{
				{
					{Value: '1', Relation: RelationNone},
					{Value: '2', Relation: RelationNone},
					{Value: '3', Relation: RelationNone},
					{Value: '4', Relation: RelationNone},
					{Value: '5', Relation: RelationNone},
				},
				{
					{Value: 'a', Relation: RelationNone},
					{Value: 'b', Relation: RelationNone},
					{Value: 'c', Relation: RelationNone},
					{Value: '9', Relation: RelationNone},
					{Value: '0', Relation: RelationNone},
				},
				{
					{Value: 'd', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
					{Value: 'f', Relation: RelationNone},
					{Value: 'd', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
				},
				{
					{Value: 'g', Relation: RelationNone},
					{Value: 'h', Relation: RelationNone},
					{Value: 'i', Relation: RelationNone},
					{Value: ' ', Relation: RelationNone},
					{Value: ' ', Relation: RelationNone},
				},
			},
			src:       []string{"12345", "67890", "abcde"},
			inputData: []string{"abc", "def", "ghi"},
			config:    PasteConfig{Y: 1, Padding: " "},
			msg:       "原点座標からの貼付け（矩形は貼り付け元の行数超過）",
		},
		{
			expect: [][]MetaRune{
				{
					{Value: '1', Relation: RelationNone},
					{Value: '2', Relation: RelationNone},
					{Value: '3', Relation: RelationNone},
					{Value: '4', Relation: RelationNone},
					{Value: '5', Relation: RelationNone},
				},
				{
					{Value: '6', Relation: RelationNone},
					{Value: 'a', Relation: RelationNone},
					{Value: 'b', Relation: RelationNone},
					{Value: 'c', Relation: RelationNone},
					{Value: '0', Relation: RelationNone},
				},
				{
					{Value: 'a', Relation: RelationNone},
					{Value: 'd', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
					{Value: 'f', Relation: RelationNone},
					{Value: 'e', Relation: RelationNone},
				},
				{
					{Value: ' ', Relation: RelationNone},
					{Value: 'g', Relation: RelationNone},
					{Value: 'h', Relation: RelationNone},
					{Value: 'i', Relation: RelationNone},
					{Value: ' ', Relation: RelationNone},
				},
			},
			src:       []string{"12345", "67890", "abcde"},
			inputData: []string{"abc", "def", "ghi"},
			config:    PasteConfig{X: 1, Y: 1, Padding: " "},
			msg:       "原点座標からの貼付け（矩形は貼り付け元の行数超過）",
		},
	}
	for i, v := range testDatas {
		got := Paste(v.src, v.inputData, v.config)
		if diff := cmp.Diff(v.expect, got); diff != "" {
			t.Error("NG", i, "\n"+diff)
		} else {
			t.Log("OK", i)
		}
	}

}
