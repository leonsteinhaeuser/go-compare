package compare

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidMatchType = errors.New("invalid match type")
	ErrValueNotANumber  = errors.New("value is not a number")
)

// MatchType defines the type of match to be performed.
type MatchType string

const (
	// MatchTypeLessThan is used to compare the response with the expected value.
	// If the response is less than the expected value the validation is successful.
	// If the response is equal or greater than the expected value the validation is not successful.
	// If the response is not a number the validation is not successful.
	MatchTypeLessThan MatchType = "lt"
	// MatchTypeLessThanOrEqual is used to compare the response with the expected value.
	// If the response is less than or equal to the expected value the validation is successful.
	// If the response is greater than the expected value the validation is not successful.
	// If the response is not a number the validation is not successful.
	MatchTypeLessThanOrEqual MatchType = "lte"
	// MatchTypeGreaterThan is used to compare the response with the expected value.
	// If the response is greater than the expected value the validation is successful.
	// If the response is equal or smaller than the expected value the validation is not successful.
	// If the response is not a number the validation is not successful.
	MatchTypeGreaterThan MatchType = "gt"
	// MatchTypeGreaterThanOrEqual is used to compare the response with the expected value.
	// If the response is greater than or equal to the expected value the validation is successful.
	// If the response is smaller than the expected value the validation is not successful.
	// If the response is not a number the validation is not successful.
	MatchTypeGreaterThanOrEqual MatchType = "gte"
	// MatchTypePercentageDeviation is used to compare the response with the expected value.
	// A percentage offset is calculated from the expected value and the response.
	// If the percentage offset is smaller than the calculated offset, the validation is successful.
	// The defined offset is added to the expected value and defines a failure tolerance.
	MatchTypePercentageDeviation MatchType = "pd"
	// MatchTypeAbsoluteOffset is used to compare the response with the expected value.
	// The defined regex is used to match the response.
	MatchTypeRegex MatchType = "re"
	// MatchTypeRange is used to compare the response with the expected value.
	// The defined range is used to match the response.
	// The value must be in between the range.
	// If the response is not a number the validation is not successful.
	MatchTypeRange MatchType = "rg"
	// MatchTypeEqual is used to compare the response with the expected value.
	// If the response is equal to the expected value the validation is successful.
	// If the response is not equal to the expected value the validation is not successful.
	MatchTypeEqual MatchType = "eq"
	// MatchTypeNotEquals is used to compare the response with the expected value.
	// If the response is not equal to the expected value the validation is successful.
	// If the response is equal to the expected value the validation is not successful.
	MatchTypeNotEquals MatchType = "neq"
	// MatchTypeEmpty is used to compare the response with the expected value.
	// If the response is empty the validation is successful.
	// If the response is not empty the validation is not successful.
	MatchTypeEmpty MatchType = "et"
	// MatchTypeNotEmpty is used to compare the response with the expected value.
	// If the response is not empty the validation is successful.
	// If the response is empty the validation is not successful.
	MatchTypeNotEmpty MatchType = "ne"
	// MatchTypeContains is used to compare the response with the expected value.
	// If the response contains the expected value the validation is successful.
	// If the response does not contain the expected value the validation is not successful.
	MatchTypeContains MatchType = "ct"
)

// Validation defines the validation specification to execute a test.
type Validation struct {
	// MatchType defines the type of the validation.
	// Possible values:
	// - lt: less than
	// - lte: less than or equal
	// - gt: greater than
	// - gte: greater than or equal
	// - pd: percentual offset
	// - re: regex
	// - rg: range
	// - eq: equals
	// - neq: not equals
	// - ne: not empty
	// - et: empty
	// - ct: contains
	MatchType MatchType
	// MatchValue defines the value operation.
	// Must only be set for "percentual offset" and "range" definitions.
	// Possible values:
	// - [0-9]-[0-9]: range definition
	// - [0-9]%: percentual offset
	// - any: regex
	MatchValue *string
	// ExpectedValue defines the expected value.
	ExpectedValue interface{}
}

// Matches validates the argument value against the validation specification.
// If the validation is successful the method returns true as validation and nil as error.
func (d Validation) Matches(value interface{}) (bool, error) {
	switch d.MatchType {
	case MatchTypeLessThan:
		val1, err := valueToInt64(d.ExpectedValue)
		if err != nil {
			return false, err
		}
		val2, err := valueToInt64(value)
		if err != nil {
			return false, err
		}
		return val2 < val1, nil
	case MatchTypeLessThanOrEqual:
		val1, err := valueToInt64(d.ExpectedValue)
		if err != nil {
			return false, err
		}
		val2, err := valueToInt64(value)
		if err != nil {
			return false, err
		}
		return val2 <= val1, nil
	case MatchTypeGreaterThan:
		val1, err := valueToInt64(d.ExpectedValue)
		if err != nil {
			return false, err
		}
		val2, err := valueToInt64(value)
		if err != nil {
			return false, err
		}
		return val2 > val1, nil
	case MatchTypeGreaterThanOrEqual:
		val1, err := valueToInt64(d.ExpectedValue)
		if err != nil {
			return false, err
		}
		val2, err := valueToInt64(value)
		if err != nil {
			return false, err
		}
		return val2 >= val1, nil
	case MatchTypePercentageDeviation:
		btsFromValue, err := json.Marshal(value)
		if err != nil {
			return false, err
		}
		btsFromExpectedValue, err := json.Marshal(d.ExpectedValue)
		if err != nil {
			return false, err
		}
		reports, err := BytesDifferent(btsFromValue, btsFromExpectedValue)
		if err != nil {
			return false, err
		}
		pcnt, err := NewPercentFromFloats(len(btsFromExpectedValue), len(reports))
		if err != nil {
			return false, err
		}
		pcntFromMatchValue, err := ParsePercentageValueFromString(*d.MatchValue)
		if err != nil {
			return false, err
		}
		return pcnt.Get() <= pcntFromMatchValue.Get(), nil
	case MatchTypeRegex:
		rp := regexp.MustCompile(*d.MatchValue)
		buf := bytes.Buffer{}
		err := gob.NewEncoder(&buf).Encode(value)
		if err != nil {
			return false, fmt.Errorf("failed to encode value: %v", err)
		}
		return rp.Match(buf.Bytes()), nil
	case MatchTypeRange:
		r := rangeParser(*d.MatchValue)
		val1, err := valueToInt64(value)
		if err != nil {
			return false, err
		}
		return val1 >= r[0] && val1 <= r[1], nil
	case MatchTypeEqual:
		return reflect.DeepEqual(d.ExpectedValue, value), nil
	case MatchTypeNotEquals:
		return !reflect.DeepEqual(d.ExpectedValue, value), nil
	case MatchTypeNotEmpty:
		if str, ok := value.(string); ok {
			if str == "" {
				return false, nil
			}
			return true, nil
		}
		return value != nil, nil
	case MatchTypeEmpty:
		if str, ok := value.(string); ok {
			if str == "" {
				return true, nil
			}
			return false, nil
		}
		return value == nil, nil
	case MatchTypeContains:
		buffer := bytes.Buffer{}
		err := gob.NewEncoder(&buffer).Encode(value)
		if err != nil {
			return false, fmt.Errorf("%w: gob envoding failed for value = %v", err, value)
		}
		return strings.Contains(buffer.String(), *d.MatchValue), nil
	default:
		return false, fmt.Errorf("%w: %s", ErrInvalidMatchType, d.MatchType)
	}
}

// valueToInt64 converts the value to an int64.
func valueToInt64(value interface{}) (int64, error) {
	iVal, ok := value.(int64)
	if !ok {
		return 0, fmt.Errorf("%w: %v", ErrValueNotANumber, value)
	}
	return iVal, nil
}

// rangeParser parses the range definition.
// The range definition is expected to be in the format:
// - [0-9]-[0-9]: range definition
func rangeParser(input string) [2]int64 {
	var result [2]int64
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return result
	}
	result[0], _ = strconv.ParseInt(parts[0], 10, 64)
	result[1], _ = strconv.ParseInt(parts[1], 10, 64)
	return result
}
