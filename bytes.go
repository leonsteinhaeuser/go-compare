package compare

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrAExceedsRangeOfB = errors.New("a must be smaller than b")
)

// bytesDifferent returns the difference between a and b.
// a must be smaller than b
func bytesDifferent(a, b []byte) (Reports, error) {
	if len(b) < len(a) {
		return nil, fmt.Errorf("%w: a=%v, b=%v", ErrAExceedsRangeOfB, a, b)
	}
	different := Reports{}
	for idx, value := range a {
		rsp := ByteDifferent(value, b[idx])
		if rsp != nil {
			rsp.Index = idx
		}
	}

	for idx := len(a); idx < len(b); idx++ {
		different = append(different, Report{
			Type:     reflect.TypeOf(b[idx]).String(),
			Index:    idx,
			Original: nil,
			New:      b[idx],
		})
	}

	return different, nil
}

// ByteDifferent returns the difference between a and b.+
// If a and b are the same, nil is returned.
func ByteDifferent(a, b byte) *Report {
	if a != b {
		return &Report{
			Type:     reflect.TypeOf(a).String(),
			Index:    0,
			Original: &a,
			New:      b,
		}
	}
	return nil
}

// BytesDifferent returns the difference between a and b.
func BytesDifferent(a, b []byte) (Reports, error) {
	deferBts, err := bytesDifferent(b, a)
	if err != nil {
		reports, err := BytesDifferent(b, a)
		if err != nil {
			return nil, err
		}
		reports.swap()
		return reports, nil
	}
	return deferBts, nil
}
