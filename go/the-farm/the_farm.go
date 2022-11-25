package thefarm

import (
    "errors"
    "fmt"
)

type SillyNephewError struct {
    cows int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}


// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	f, err := weightFodder.FodderAmount()

    if f < 0 && err == nil {
        return 0, errors.New("negative fodder")
    }

    if cows < 0 {
        return 0, &SillyNephewError{
            cows: cows,
        }
    }

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if err != nil {
		if err == ErrScaleMalfunction {
            if f > 0 {
                return (f * 2.0) / float64(cows), nil
            }

            return 0, errors.New("negative fodder")
		}

		return 0, err
	}


	return f / float64(cows), nil
}
