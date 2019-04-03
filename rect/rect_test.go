package rect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaste(t *testing.T) {
	assert.Equal(t, []string{}, Paste([]string{"12345","67890","abcde"}, []string{"abc","def", "efg"}, PasteConfig{}))
}
