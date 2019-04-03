package rect

import (
	"testing"

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
	assert.Equal(t, "abc  ", PadSpace("abcde", "abc", PasteConfig{X: 0, Padding: " "}))
	assert.Equal(t, " abc ", PadSpace("abcde", "abc", PasteConfig{X: 1, Padding: " "}))
	assert.Equal(t, "  abc", PadSpace("abcde", "abc", PasteConfig{X: 2, Padding: " "}))
	assert.Equal(t, "   abc", PadSpace("abcde", "abc", PasteConfig{X: 3, Padding: " "}))

	assert.Equal(t, "abc 　", PadSpace("abcdef", "abc", PasteConfig{X: 0, Padding: "　"}))
	assert.Equal(t, "abc　", PadSpace("abcde", "abc", PasteConfig{X: 0, Padding: "　"}))
	assert.Equal(t, " abc ", PadSpace("abcde", "abc", PasteConfig{X: 1, Padding: "　"}))
	assert.Equal(t, "　abc", PadSpace("abcde", "abc", PasteConfig{X: 2, Padding: "　"}))
	assert.Equal(t, "　 abc", PadSpace("abcde", "abc", PasteConfig{X: 3, Padding: "　"}))
}

// func TestPaste(t *testing.T) {
// 	assert.Equal(t, []string{"abc45", "def90", "ghide"}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: " "}), "Paste 0,0 axis")
// 	assert.Equal(t, []string{"12345", "6abc0", "adefe", " ghi "}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: " "}), "Paste 1,1 axis")
// 	assert.Equal(t, []string{"123456", "6abc01", "adefef", " ghi"}, Paste([]string{"123456", "678901", "abcdef"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: "　"}), "Paste with fullwidth space")
// 	assert.Equal(t, []string{"abc456", "def901", "ghi　　"}, Paste([]string{"123456", "678901"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: "　"}), "Paste with fullwidth space")
// 	assert.Equal(t, []string{"あい5", "うえ0"}, Paste([]string{"12345", "67890"}, []string{"あい", "うえ"}, PasteConfig{}), "Paste fullwidth character")
// 	assert.Equal(t, []string{" あい ", " うえ "}, Paste([]string{"１２３", "４５６"}, []string{"あい", "うえ"}, PasteConfig{X: 1}), "Paste fullwidth character")
// }
