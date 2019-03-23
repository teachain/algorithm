package algorithm

func findSmallest(arr []int)int{
	smallest:=arr[0]//存储最小的值
	index:=0 //存储最小元素的索引
	for i:=range arr{
		if arr[i]<smallest{
			smallest=arr[i]
			index=i
		}
	}
	return index
}

//选择排序
//它的基本原则就是每次找出来的元素的位置肯定是确定的
//它每次是从剩余元素中找到最值
func SelectionSort(arr []int)[]int{ //对数组进行排序
	newArr:=make([]int,0,len(arr))
	for len(arr)>0{
		smallest:=findSmallest(arr) //找出数组中最小的元素
		newArr=append(newArr,arr[smallest])//将其加入到新数组中
		arr=append(arr[:smallest],arr[smallest+1:]...)
	}
	return newArr
}
