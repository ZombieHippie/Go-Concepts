package GoAlgorithmsMIT

import (
	"math/rand"
	"testing"
	"time"
)

type list []int64

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l list) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l list) Len() int {
	return len(l)
}

func createList(size int) list {
	l := make(list, size)
	for i := range l {
		l[i] = int64(i)
	}
	return l
}

func TestInsertionSortSorted(t *testing.T) {
	testListLength := 12
	testList := createList(testListLength)
	InsertionSort(testList, 0, testListLength-1)
	for i := 1; i < testListLength; i++ {
		if testList[i-1] > testList[i] {
			t.Log(testList)
			t.FailNow()
		}
	}
}
func TestInsertionSortReverseSorted(t *testing.T) {
	testListLength := 12
	testList := createList(testListLength)
	for i := 0; i < testListLength; i++ {
		testList[i] = int64(testListLength - i)
	}
	lastIndice := testListLength - 1
	firstIndice := 0
	t.Log(testList, firstIndice, lastIndice)
	InsertionSort(testList, firstIndice, lastIndice)
	for i := 1; i < testListLength; i++ {
		if testList[i-1] > testList[i] {
			t.Log(testList)
			t.FailNow()
		}
	}
}

func BenchmarkInsertionSortE3(b *testing.B) {
	rand.Seed(int64(time.Now().Nanosecond()))
	benches := b.N // number of random lists
	listLength := 1000
	listsToSort := make([]list, benches)
	for i := range listsToSort {
		listsToSort[i] = createList(listLength)
		for j := range listsToSort[i] {
			listsToSort[i][j] = rand.Int63()
		}
	}
	// start timer
	lastIndice := listLength - 1
	b.ResetTimer()
	for _, lst := range listsToSort {
		InsertionSort(lst, 0, lastIndice)
	}
}
