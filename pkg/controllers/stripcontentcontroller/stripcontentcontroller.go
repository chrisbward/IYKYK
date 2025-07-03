package stripcontentcontroller

import (
	"errors"
	"regexp"
	"strings"

	"github.com/chrisbward/IYKYK/pkg/entities"
	"github.com/chrisbward/IYKYK/pkg/util"
	"github.com/forPelevin/gomoji"
)

type StripContentController struct {
}

// NewContentController returns an instance of the controller
func NewStripContentController() (stripContentController entities.IStripContentController, err error) {

	stripContentController = &StripContentController{}

	return
}

// StripInputOfEmoji removes all emoji characters from the input string using the gomoji library.
//
// Parameters:
//   - input: The input string potentially containing emojis.
//
// Returns:
//   - output: The cleaned string with all emojis removed.
//   - err: An error if the input is empty; otherwise, nil.
func (scc *StripContentController) StripInputOfEmoji(input string) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}

	output = util.CollpaseWhitespace(gomoji.RemoveEmojis(input))
	return
}

// StripInputOfEmDash replaces all em dashes ("—") in the input string with standard hyphens ("-").
// It also normalizes spacing around the hyphens to ensure a single space on each side.
//
// Parameters:
//   - input: The input string to process.
//
// Returns:
//   - output: The resulting string with em dashes replaced and spacing normalized.
//   - err: An error if the input is empty; otherwise, nil.
func (scc *StripContentController) StripInputOfEmDash(input string) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}

	re := regexp.MustCompile(`—`)
	output = re.ReplaceAllString(input, " - ")

	// Finally, normalize multiple spaces to a single space
	output = util.CollpaseWhitespace(output)

	return
}

// StripInputOfAngledQuotes replaces curly or angled quotation marks in the input string
// with standard ASCII quotes (i.e., ' and ").
//
// Parameters:
//   - input: The input string to be cleaned.
//
// Returns:
//   - output: The cleaned string with standard quotes.
//   - err: An error if the input is empty; otherwise, nil.
func (scc *StripContentController) StripInputOfAngledQuotes(input string) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}
	replacer := strings.NewReplacer(
		"’", "'",
		"‘", "'",
		"“", `"`,
		"”", `"`,
	)

	output, err = scc.ReplaceInput(input, replacer)
	if err != nil {
		return "", err
	}
	return
}

// ReplaceInput replaces substrings in the input string using the provided strings.Replacer.
//
// Parameters:
//   - input: The input string to be modified.
//   - replacer: A strings.Replacer used to define substitution rules.
//
// Returns:
//   - output: The resulting string after replacements.
//   - err: An error if the input is empty; otherwise, nil.
func (scc *StripContentController) ReplaceInput(input string, replacer *strings.Replacer) (output string, err error) {
	if input == "" {
		err = errors.New("input is empty")
		return "", err
	}
	output = replacer.Replace(input)
	return
}
