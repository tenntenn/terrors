//go:build go1.23

package terrors

import (
	"errors"
	"iter"
)

// All creates an iterator of all unwrapped errors, including joined errors.
func All(err error) iter.Seq[error] {
	return func(yield func(error) bool) {
		all(err, yield)
	}
}

// AsAll creates an iterator of type T values that are converted from the each errors when errors.As returns true.
// If an error is root of joined errors, it will be skipped.
func AsAll[T any](err error) iter.Seq[T] {
	return func(yield func(T) bool) {
		all(err, func(err error) bool {
			// skip a joined error
			_, joined := err.(interface{ Unwrap() []error })
			if joined {
				return true
			}

			var v T
			if !errors.As(err, &v) {
				return true
			}

			if !yield(v) {
				return false
			}

			return true
		})
	}
}

func all(err error, yield func(error) bool) bool {
	if err == nil {
		return true
	}

	if !yield(err) {
		return false
	}

	if wrapped := errors.Unwrap(err); wrapped != nil {
		return all(wrapped, yield)
	}

	joined, ok := err.(interface{ Unwrap() []error })
	if !ok {
		return true
	}

	for _, child := range joined.Unwrap() {
		if !all(child, yield) {
			return false
		}
	}

	return true
}
