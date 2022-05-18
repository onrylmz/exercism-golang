package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	nephew int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.nephew)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}

	if err == ErrScaleMalfunction {
		fodder *= 2
	}

	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("Division by zero")
	}

	if cows < 0 {
		return 0, &SillyNephewError{nephew: cows}
	}

	return fodder / float64(cows), nil
}
