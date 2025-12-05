package utils

import "golang.org/x/exp/constraints"

func Mod[T constraints.Integer](n, m T) T {
	if m == 0 {
		panic("modulo by zero")
	}

	r := n % m
	if r < 0 {
		r += m
	}
	return r
}

func Abs[T constraints.Integer | constraints.Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Sum[T constraints.Integer | constraints.Float](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Product[T constraints.Integer | constraints.Float](nums []T) T {
	if len(nums) == 0 {
		return 0
	}
	product := T(1)
	for _, n := range nums {
		product *= n
	}
	return product
}
