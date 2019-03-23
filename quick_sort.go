package algorithm

//快速排序使用了分而治之的方法

func QuickSort(arr []int)[]int{
	result:=make([]int,0)
	if len(arr)<2{//基线条件：为空或只包含一个元素的数组是"有序"的
		return arr
	}else{//递归条件
	    less:=make([]int,0)
		greater:=make([]int,0)
		pivot:=arr[0]
		for i:=1;i<len(arr);i++{
			if arr[i]<=pivot{
				less=append(less,arr[i])
			}else{
				greater=append(greater,arr[i])
			}
		}
		lessResult:=QuickSort(less)
		if len(lessResult)>0{
			result=append(result,lessResult...)
		}
		result=append(result,pivot)
		greaterResult:=QuickSort(greater)
		if len(greaterResult)>0{
			result=append(result,greaterResult...)
		}
		return result
	}
}
