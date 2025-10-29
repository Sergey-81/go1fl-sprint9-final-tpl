package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random positive integers
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		data[i] = r.Int()
	}

	return data
}

// maximum returns the maximum number in the slice
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number using concurrent processing
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) <= CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(data)
		}
		
		chunk := data[startIndex:endIndex]
		
		go func(chunkIndex int, chunk []int) {
			defer wg.Done()
			maxValues[chunkIndex] = maximum(chunk)
		}(i, chunk)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	
	start := time.Now()
	data := generateRandomElements(SIZE)
	generationTime := time.Since(start)
	fmt.Printf("Время генерации: %d μs\n", generationTime.Microseconds())

	fmt.Println("Ищем максимальное значение в один поток")
	start = time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", maxSingle, elapsedSingle.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", maxMulti, elapsedMulti.Microseconds())
	
	if maxSingle != maxMulti {
		fmt.Println("ВНИМАНИЕ: Результаты не совпадают!")
	} else {
		fmt.Println("Результаты совпадают ✓")
	}
}
