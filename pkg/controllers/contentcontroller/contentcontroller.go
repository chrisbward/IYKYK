package contentcontroller

import (
	"errors"

	"github.com/chrisbward/IYKYK/pkg/entities"
)

// Options holds configuration flags that control which cleaning operations
// should be applied automatically by ContentController methods.
//
// Fields:
//   - StripEmoji: When true, emoji characters will be stripped from input.
//   - StripEmDash: When true, em dashes will be replaced by hyphens.
//   - StripAngledQuotes: When true, angled or curly quotes will be normalized to ASCII quotes.
type ContentControllerOptions struct {
	StripEmDash       bool
	StripEmoji        bool
	StripAngledQuotes bool
}

type ContentController struct {
	Options                ContentControllerOptions
	StripContentController entities.IStripContentController
}

// NewContentController returns an instance of the controller
func NewContentController(options *ContentControllerOptions, stripContentController entities.IStripContentController) (contentController entities.IContentController, err error) {

	// set defaults if options are not passed
	if options == nil {
		options = &ContentControllerOptions{
			StripEmDash:       false,
			StripEmoji:        false,
			StripAngledQuotes: false,
		}
	}

	contentController = &ContentController{
		Options:                *options,
		StripContentController: stripContentController,
	}

	return
}

// CleanContentAutomatic applies a series of cleaning operations to the input string
// based on the controller's configuration options (e.g., strip emojis, em dashes, or angled quotes).
//
// Parameters:
//   - input: The input string to be cleaned.
//
// Returns:
//   - output: The cleaned string after applying all enabled transformations.
//   - err: An error if the input is empty or if any cleaning operation fails.
func (cc *ContentController) CleanContentAutomatic(input string) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}

	stripFuncs := make([]entities.StripFunction, 0)

	if cc.Options.StripEmoji {
		stripFuncs = append(stripFuncs, cc.StripContentController.StripInputOfEmoji)
	}
	if cc.Options.StripEmDash {
		stripFuncs = append(stripFuncs, cc.StripContentController.StripInputOfEmDash)
	}
	if cc.Options.StripAngledQuotes {
		stripFuncs = append(stripFuncs, cc.StripContentController.StripInputOfAngledQuotes)
	}

	output, err = cc.DoCleanWithStripFunctions(input, stripFuncs...)
	return
}

// DoCleanWithStripFunctions applies a sequence of stripping functions to the input string,
// in the order they are provided. If any function returns an error, the process stops immediately.
//
// Parameters:
//   - input: The input string to be processed.
//   - funcs: A variadic list of StripFunction functions to apply.
//
// Returns:
//   - output: The resulting string after all stripping functions have been applied successfully.
//   - err: An error if the input is empty, no functions are provided, or any function fails during processing.
func (cc *ContentController) DoCleanWithStripFunctions(input string, funcs ...entities.StripFunction) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}
	if len(funcs) == 0 {
		err = errors.New("must provide strip functions")
		return "", err
	}

	output = input
	for _, stripFunc := range funcs {
		output, err = stripFunc(output)
		if err != nil {
			return // stop and return error immediately
		}
	}
	return
}
