package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "обычный размер", size: 100, want: 100},
		{name: "один элемент", size: 1, want: 1},
		{name: "нулевой размер", size: 0, want: 0},
		{name: "отрицательный размер", size: -5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.want {
				t.Fatalf("длина = %d, ожидается %d", len(data), tt.want)
			}

			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d: число должно быть положительным",
						i, value)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "максимум в начале",
			data: []int{99, 5, 3},
			want: 99,
		},
		{
			name: "максимум в середине",
			data: []int{3, 99, 5},
			want: 99,
		},
		{
			name: "максимум в конце",
			data: []int{3, 5, 99},
			want: 99,
		},
		{
			name: "одинаковые элементы",
			data: []int{7, 7, 7},
			want: 7,
		},
		{
			name: "ровно восемь элементов",
			data: []int{1, 2, 3, 4, 50, 6, 7, 8},
			want: 50,
		},
		{
			name: "максимум в остатке",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100},
			want: 100,
		},
		{
			name: "длина кратна восьми",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8,
				9, 10, 11, 12, 13, 14, 99, 16},
			want: 99,
		},
		{
			name: "отрицательные элементы",
			data: []int{-10, -3, -7},
			want: -3,
		},
	}

	functions := []struct {
		name string
		find func([]int) int
	}{
		{name: "maximum", find: maximum},
		{name: "maxChunks", find: maxChunks},
	}

	for _, fn := range functions {
		t.Run(fn.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := fn.find(tt.data)

					if got != tt.want {
						t.Errorf("получено %d, ожидается %d", got, tt.want)
					}
				})
			}
		})
	}
}
