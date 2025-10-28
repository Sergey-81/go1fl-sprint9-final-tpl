package main

import "testing"

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
			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = length %d, expected %d", 
					tt.size, len(result), tt.expected)
			}
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, expected %d", tt.data, result, tt.expected)
			}
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
		{"Multiple elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, expected %d", tt.data, result, tt.expected)
			}
		})
	}
}