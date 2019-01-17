package sort_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	mysort "github.com/x899/algorithms/sort"
)

func testSort(t *testing.T, sortfunc func([]int)) {
	tests := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, test := range tests {
		dataLen := len(test.original)
		sortfunc(test.original)

		for i := 0; i < dataLen; i++ {
			if test.original[i] != test.sorted[i] {
				t.Errorf("wrong output expected: %v got: %v",
					test.sorted, test.original)
			}
		}
	}
}

func testSortRandom(t *testing.T, sortfunc func([]int)) {
	for i := 0; i < 5; i++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		array1 := make([]int, random.Intn(10000))
		for i := range array1 {
			array1[i] = random.Intn(10000)
		}
		array2 := make([]int, len(array1))
		copy(array2, array1)

		sortfunc(array1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("wrong output expected: %v got: %v",
					array2, array1)
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	testSort(t, mysort.InsertionSort)
	testSortRandom(t, mysort.InsertionSort)
}

func TestBubbleSort(t *testing.T) {
	testSort(t, mysort.BubbleSort)
	testSortRandom(t, mysort.BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	testSort(t, mysort.SelectionSort)
	testSortRandom(t, mysort.SelectionSort)
}

func TestShellSort(t *testing.T) {
	testSort(t, mysort.ShellSort)
	testSortRandom(t, mysort.ShellSort)
}

func TestMergeSortBottomUp(t *testing.T) {
	testSort(t, mysort.MergeSortBottomUp)
	testSortRandom(t, mysort.MergeSortBottomUp)
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, test := range tests {
		dataLen := len(test.original)
		mysort.MergeSort(test.original, 0, dataLen-1)

		for i := 0; i < dataLen; i++ {
			if test.original[i] != test.sorted[i] {
				t.Errorf("wrong output expected: %v got: %v",
					test.sorted, test.original)
			}
		}
	}
}

func TestMergeSortRandom(t *testing.T) {
	for i := 0; i < 5; i++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		array1 := make([]int, random.Intn(10000))
		for i := range array1 {
			array1[i] = random.Intn(10000)
		}
		array2 := make([]int, len(array1))
		copy(array2, array1)

		mysort.MergeSort(array1, 0, len(array1)-1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("wrong output expected: %v got: %v",
					array2, array1)
			}
		}
	}
}

func TestMergeSortOptimized(t *testing.T) {
	tests := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, test := range tests {
		dataLen := len(test.original)
		aux := make([]int, dataLen)
		mysort.MergeSortOptimized(test.original, aux, 0, dataLen-1)

		for i := 0; i < dataLen; i++ {
			if test.original[i] != test.sorted[i] {
				t.Errorf("wrong output expected: %v got: %v",
					test.sorted, test.original)
			}
		}
	}
}

func TestMergeSortOptimizedRandom(t *testing.T) {
	for i := 0; i < 5; i++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		array1 := make([]int, random.Intn(10000))
		for i := range array1 {
			array1[i] = random.Intn(10000)
		}
		array2 := make([]int, len(array1))
		copy(array2, array1)

		aux := make([]int, len(array1))
		mysort.MergeSortOptimized(array1, aux, 0, len(array1)-1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("wrong output expected: %v got: %v",
					array2, array1)
			}
		}
	}
}