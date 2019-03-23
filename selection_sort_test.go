package algorithm

import "testing"

func TestSelectionSort(t *testing.T) {
	arr:=[]int{5,3,6,2,10,98,78,65}
	result:=SelectionSort(arr)
	t.Log("result",result)
}
