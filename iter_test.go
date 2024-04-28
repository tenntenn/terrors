//go:build go1.23

package terrors_test

import (
	"errors"
	"fmt"
	"iter"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/tenntenn/terrors"
)

func TestAll(t *testing.T) {
	t.Parallel()

	errA := errors.New("A")
	errB := errors.New("B")
	errAB := errors.Join(errA, errB)
	errC := fmt.Errorf("C: %w", errA)
	errBC := errors.Join(errB, errC)

	cases := []struct {
		name string
		err  error
		want []error
	}{
		{"single", errA, errs(errA)},
		{"joined", errAB, errs(errAB, errA, errB)},
		{"wrapped", errC, errs(errC, errA)},
		{"wrapped and joined", errBC, errs(errBC, errB, errC, errA)},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := toSlice(terrors.All(tt.err))
			if diff := cmp.Diff(got, tt.want, cmpopts.EquateErrors()); diff != "" {
				t.Error("(got, want)", diff)
			}
		})
	}
}

func errs(err ...error) []error {
	return err
}

func toSlice[T any](seq iter.Seq[T]) []T {
	var s []T
	for v := range seq {
		s = append(s, v)
	}
	return s
}
