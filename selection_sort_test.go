package algorithm

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 6, 2, 10, 98, 78, 65}
	result := SelectionSort(arr)
	fmt.Println("result", result)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 6, 2, 10, 98, 78, 65}
	BubbleSort(arr)
	fmt.Println("TestBubbleSort result", arr)
}
