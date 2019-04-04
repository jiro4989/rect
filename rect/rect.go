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

func PasteLine(src, inputData string, config PasteConfig) (ret string) {
	return
}
