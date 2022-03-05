package compare

import (
	"errors"
	"fmt"
)

var (
	// ErrValueExceedsRange is returned when the value exceeds the supported range.
	ErrValueExceedsRange = errors.New("value exceeds supported range")

	// ErrAorBIsZero is returned when a or b is 0.
	ErrAorBIsZero = errors.New("a or b is 0")
)

// Percent represents a custom type that defines a specific allowed value range from 0.0 to 100.0.
type Percent struct {
	value float64
}

// NewPercent returns a new instance of type Percent.
func NewPercent(value float64) (*Percent, error) {
	pcnt := &Percent{}
	err := pcnt.Set(value)
	if err != nil {
		return nil, err
	}
	return pcnt, nil
}

// checkForRangeCondition checks if the value is within the supported range.
func (p *Percent) checkForRangeCondition(value float64) error {
	const (
		upperBorder = 100.0
		lowerBorder = 0.0
	)
	if value > upperBorder || value < lowerBorder {
		return fmt.Errorf("%w: value=%v, supported range: 0-100", ErrValueExceedsRange, value)
	}
	return nil
}

// Set sets the value of the Percent type.
func (p *Percent) Set(value float64) error {
	if err := p.checkForRangeCondition(value); err != nil {
		return err
	}
	p.value = value
	return nil
}

// Get returns the value of the Percent type.
// If the value is not within the supported range, 0.0 is returned.
func (p Percent) Get() float64 {
	if err := p.checkForRangeCondition(p.value); err != nil {
		return 0
	}
	return p.value
}

// NewPercentFromInts calculates the percent difference between a and b.
func NewPercentFromFloats(a, b int) (*Percent, error) {
	aFloat := float64(a)
	bFloat := float64(b)

	if a == 0 && b == 0 {
		return nil, fmt.Errorf("%w: a=%v, b=%v", ErrAorBIsZero, a, b)
	}

	if a == b {
		return NewPercent(0.0)
	}

	var p float64
	if aFloat > bFloat {
		p = (bFloat / aFloat) * 100
	} else {
		p = (aFloat / bFloat) * 100
	}

	return NewPercent(p)
}
