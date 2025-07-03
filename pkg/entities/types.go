package entities

// StripFunction defines the signature for functions that process an input string
// and return a cleaned output string along with an error if processing fails.
//
// It is typically used to represent individual string cleaning operations.
type StripFunction func(input string) (output string, err error)
