package rect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaste(t *testing.T) {
	assert.Equal(t, []string{"abc45", "def90", "ghide"}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: " "}), "Paste 0,0 axis")
	assert.Equal(t, []string{"12345", "6abc0", "adefe", " ghi "}, Paste([]string{"12345", "67890", "abcde"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: " "}), "Paste 1,1 axis")
	// assert.Equal(t, []string{"123456", "6abc01", "adefef", " ghi 　"}, Paste([]string{"123456", "678901", "abcdef"}, []string{"abc", "def", "ghi"}, PasteConfig{X: 1, Y: 1, Padding: "　"}), "Paste with fullwidth space")
	// assert.Equal(t, []string{"abc456", "def901", "ghi　　"}, Paste([]string{"123456", "678901"}, []string{"abc", "def", "ghi"}, PasteConfig{Padding: "　"}), "Paste with fullwidth space")
}
