package sort_test

import (
	"homework/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sortTestSet struct {
	title  string
	input  []int
	expect []int
}

var testSets = []sortTestSet{
	// Test set 1
	{
		title:  "Test sort: 1, 5, 4, 3, 2",
		input:  []int{1, 5, 4, 3, 2},
		expect: []int{1, 2, 3, 4, 5},
	},
	{
		title:  "Test sort: -5, -4, -100, 9, 0",
		input:  []int{-5, -4, -100, 9, 0},
		expect: []int{-100, -5, -4, 0, 9},
	},
	{
		title:  "Test sort: 1",
		input:  []int{1},
		expect: []int{1},
	},
	{
		title:  "Test sort: 3, 2, 0",
		input:  []int{3, 2, 0},
		expect: []int{0, 2, 3},
	},
	{
		title:  "Test sort: 10, 10, 10, 10",
		input:  []int{10, 10, 10, 10},
		expect: []int{10, 10, 10, 10},
	},
	{
		title:  "Test sort: nil",
		input:  nil,
		expect: nil,
	},
}

func TestSort(t *testing.T) {
	t.Run("BubbleSort", func(t *testing.T) {
		for _, ts := range testSets {
			t.Run(ts.title, func(t *testing.T) {
				sort.BubbleSort(ts.input)
				assert.Equal(t, ts.expect, ts.input)
			})
		}
	})
}
