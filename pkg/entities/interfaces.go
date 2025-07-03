package entities

import "strings"

type IContentController interface {
	CleanContentAutomatic(input string) (output string, err error)
	DoCleanWithStripFunctions(input string, funcs ...StripFunction) (output string, err error)
}

type IStripContentController interface {
	StripInputOfEmoji(input string) (output string, err error)
	StripInputOfEmDash(input string) (output string, err error)
	StripInputOfAngledQuotes(input string) (output string, err error)
	ReplaceInput(input string, replacer *strings.Replacer) (output string, err error)
}
