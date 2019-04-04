package rect

type PasteConfig struct {
	X                int
	Y                int
	Padding          string
	UseInsert        bool
	IgnoreWhiteSpace bool
}

func Paste(src, inputData []string, config PasteConfig) (ret []string) {
	return
}

func PasteLine(src, inputData string, config PasteConfig) string {
	srcRune := []rune(src)
	inputRune := []rune(inputData)
	for i := 0; i < len(inputRune); i++ {
		ir := inputRune[i]
		i2 := i + config.X
		if len(srcRune) <= i2 {
			srcRune = append(srcRune, ' ')
		}
		srcRune[i2] = ir
	}
	return string(srcRune)
}
