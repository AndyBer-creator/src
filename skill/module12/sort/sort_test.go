package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {

	b.Run("small arrays", func(b *testing.B) {
		testArr := generateSlice(10, 10)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			BubbleSort(ar)
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		testArr := generateSlice(100, 1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			BubbleSort(ar)
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		testArr := generateSlice(10000, 100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			BubbleSort(ar)
		}
	})

}
func BenchmarkSelectionSort(b *testing.B) {

	b.Run("small arrays", func(b *testing.B) {
		testArr := generateSlice(10, 10)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			SelectionSort(ar)
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		testArr := generateSlice(100, 1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			SelectionSort(ar)
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		testArr := generateSlice(10000, 100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			SelectionSort(ar)
		}
	})
}

func BenchmarkInsertionSort(b *testing.B) {

	b.Run("small arrays", func(b *testing.B) {
		testArr := generateSlice(10, 10)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			InsertionSort(ar)
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		testArr := generateSlice(100, 1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			InsertionSort(ar)
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		testArr := generateSlice(10000, 100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			InsertionSort(ar)
		}
	})
}

func BenchmarkMergeSort(b *testing.B) {

	b.Run("small arrays", func(b *testing.B) {
		testArr := generateSlice(10, 10)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			MergeSort(ar)
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		testArr := generateSlice(100, 1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			MergeSort(ar)
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		testArr := generateSlice(10000, 100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			MergeSort(ar)
		}
	})
}

func BenchmarkQuickSort(b *testing.B) {

	b.Run("small arrays", func(b *testing.B) {
		testArr := generateSlice(10, 10)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			QuickSort(ar, 0, len(ar)-1)
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		testArr := generateSlice(100, 1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			QuickSort(ar, 0, len(ar)-1)
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		testArr := generateSlice(10000, 100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := make([]int, len(testArr))
			copy(ar, testArr)
			QuickSort(ar, 0, len(ar)-1)
		}
	})
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
