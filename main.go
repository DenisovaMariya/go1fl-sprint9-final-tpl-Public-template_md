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

// generateRandomElements создаёт слайс положительных случайных чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum возвращает максимум или 0 для пустого слайса.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks находит максимум с помощью CHUNKS горутин.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Последняя часть включает остаток от деления.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()

			if len(chunk) == 0 {
				// Пустые части не должны влиять на общий максимум.
				maxValues[index] = data[0]
				return
			}

			maxValues[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		maxValue,
		elapsed,
	)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)

	start = time.Now()
	parallelMax := maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		parallelMax,
		elapsed,
	)

	if maxValue != parallelMax {
		fmt.Println("Ошибка: результаты поиска не совпадают")
	}
}
