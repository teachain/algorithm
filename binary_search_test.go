package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
    list:=[]int {1,3,5,7,9}
    pos:=BinarySearch(list,7)
    if pos>=0{
    	t.Log("found it",7)
	}else{
		t.Log("not found it",7)
	}
    pos=BinarySearch(list,11)
	if pos>=0{
		t.Log("found it",11)
	}else{
		t.Log("not found it",11)
	}
    result:=math.Log2(1000000000)

	t.Log("result",result)
}

func TestBinarySearch2(t *testing.T) {
	countDown(100)
}

func countDown(i int){
	fmt.Println(i)
	if i<=0{  //基线条件
		return
	}else{ //递归条件
		countDown(i-1)
	}
}
