package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Zero size", 0, 0},
		{"Small size", 10, 10},
		{"Medium size", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			assert.Equal(t, tt.expected, len(result), "generateRandomElements(%d) should return slice of length %d", tt.size, tt.expected)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 5, 3, 9, 2}, 9},
		{"All same elements", []int{7, 7, 7, 7}, 7},
		{"Negative numbers", []int{-5, -1, -10, -3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			assert.Equal(t, tt.expected, result, "maximum(%v) should return %d", tt.data, tt.expected)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Less elements than chunks", []int{1, 2, 3}, 3},
		{"Exactly chunks elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"Large dataset", make([]int, 1000), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For large dataset, initialize with values
			if tt.name == "Large dataset" {
				for i := range tt.data {
					tt.data[i] = i
				}
				tt.expected = 999
			}
			
			result := maxChunks(tt.data)
			assert.Equal(t, tt.expected, result, "maxChunks(%v) should return %d", tt.data, tt.expected)
		})
	}
}

func TestConsistencyBetweenSingleAndMultiThreaded(t *testing.T) {
	data := generateRandomElements(1000)
	
	maxSingle := maximum(data)
	maxMulti := maxChunks(data)
	
	assert.Equal(t, maxSingle, maxMulti, "Single-threaded and multi-threaded results should be equal")
}

func TestMaxChunksWithVariousSizes(t *testing.T) {
	testCases := []struct {
		size     int
	}{
		{0},
		{1},
		{7},  // less than CHUNKS
		{8},  // equal to CHUNKS
		{16}, // multiple of CHUNKS
		{20}, // not multiple of CHUNKS
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Size%d", tc.size), func(t *testing.T) {
			data := make([]int, tc.size)
			for i := range data {
				data[i] = i
			}
			if tc.size > 0 {
				data[tc.size-1] = tc.size * 2 // Set max value
			}

			expected := 0
			if tc.size > 0 {
				expected = tc.size * 2
			}

			result := maxChunks(data)
			assert.Equal(t, expected, result, "For size %d", tc.size)
		})
	}
}