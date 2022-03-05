package compare

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidPercentageFormat = fmt.Errorf("invalid percentage format")
)

// ParsePercentageValueFromString parses a percentage value from a string and returns the value as type Percent.
// If the string is not a valid percentage value, an error is returned.
func ParsePercentageValueFromString(in string) (*Percent, error) {
	if !strings.HasSuffix(in, "%") {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPercentageFormat, in)
	}
	value, err := strconv.ParseFloat(strings.TrimSuffix(in, "%"), 64)
	if err != nil {
		return nil, err
	}
	return NewPercent(value)
}
