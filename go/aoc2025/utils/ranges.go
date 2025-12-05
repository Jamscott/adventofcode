package utils

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Range[T Numeric] struct {
	Start T
	End   T
}

func (r Range[T]) Contains(value T) bool {
	return r.Start <= value && value <= r.End
}

func (r Range[T]) Overlaps(other Range[T]) bool {
	return r.Start <= other.End && other.Start <= r.End
}

func (r Range[T]) Length() T {
	return r.End - r.Start + 1
}

type RangeSet[T Numeric] []Range[T]

func (rs RangeSet[T]) Contains(value T) bool {
	for _, r := range rs {
		if r.Contains(value) {
			return true
		}
	}
	return false
}

func (rs RangeSet[T]) Sort() {
	slices.SortFunc(rs, func(a, b Range[T]) int {
		return cmp.Compare(a.Start, b.Start)
	})
}

func ParseInt64Range(s string) (Range[int64], error) {
	parts := strings.SplitN(s, "-", 2)
	if len(parts) != 2 {
		return Range[int64]{}, strconv.ErrSyntax
	}

	start, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	if err != nil {
		return Range[int64]{}, err
	}

	end, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
	if err != nil {
		return Range[int64]{}, err
	}

	return Range[int64]{Start: start, End: end}, nil
}
