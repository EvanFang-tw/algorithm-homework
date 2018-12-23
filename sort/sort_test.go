package sort_test

import (
	"homework/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort1(t *testing.T) {
	input := []int{1, 5, 4, 3, 2}

	expect := []int{1, 2, 3, 4, 5}

	sort.BubbleSort(input)

	assert.Equal(t, expect, input)
}

func TestBubbleSort2(t *testing.T) {
	input := []int{5, -7, 4}

	expect := []int{-7, 4, 5}

	sort.BubbleSort(input)

	assert.Equal(t, expect, input)
}

func TestBubbleSort3(t *testing.T) {
	input := []int{1}

	expect := []int{1}

	sort.BubbleSort(input)

	assert.Equal(t, expect, input)
}

func TestBubbleSort4(t *testing.T) {
	input := []int{1000, 1001, 1002, 1003}

	expect := []int{1000, 1001, 1002, 1003}

	sort.BubbleSort(input)

	assert.Equal(t, expect, input)
}

func TestBubbleSort5(t *testing.T) {
	var input []int
	input = nil

	var expect []int
	expect = nil

	sort.BubbleSort(input)

	assert.Equal(t, expect, input)
}
