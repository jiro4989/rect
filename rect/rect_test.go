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

// func TestPaste(t *testing.T) {
// 	assert.Equal(t, []string{"abc45", "def90", "ghide"}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: " "}), "Paste 0,0 axis")
// 	assert.Equal(t, []string{"12345", "6abc0", "adefe", " ghi "}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: " "}), "Paste 1,1 axis")
// 	assert.Equal(t, []string{"123456", "6abc01", "adefef", " ghi"}, Paste([]string{"123456", "678901", "abcdef"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: "　"}), "Paste with fullwidth space")
// 	assert.Equal(t, []string{"abc456", "def901", "ghi　　"}, Paste([]string{"123456", "678901"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: "　"}), "Paste with fullwidth space")
// 	assert.Equal(t, []string{"あい5", "うえ0"}, Paste([]string{"12345", "67890"}, []string{"あい", "うえ"}, PasteConfig{}), "Paste fullwidth character")
// 	assert.Equal(t, []string{" あい ", " うえ "}, Paste([]string{"１２３", "４５６"}, []string{"あい", "うえ"}, PasteConfig{X: 1}), "Paste fullwidth character")
// }
