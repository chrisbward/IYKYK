package util

import (
	"regexp"
	"strings"
)

func CollpaseWhitespace(s string) string {

	// Replace any remaining non-breaking spaces explicitly
	s = strings.ReplaceAll(s, "\u00A0", " ")

	// Collapse all kinds of whitespace into a single space
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	// remove spaces before punctuation
	spaceBeforePunctRegex := regexp.MustCompile(`\s+([.,!?;:])`)
	s = spaceBeforePunctRegex.ReplaceAllString(s, "$1")

	return strings.TrimSpace(s)
}
