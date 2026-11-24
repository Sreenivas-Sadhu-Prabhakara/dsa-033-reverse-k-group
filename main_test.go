package main

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}}, // group shorter than k stays
		{[]int{}, 3, []int{}},
	}
	for _, c := range cases {
		got := toSlice(ReverseKGroup(fromSlice(c.in), c.k))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ReverseKGroup(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
