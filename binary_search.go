package algorithm

// 二分查找的运行时间为对数时间
// 也就是math.log2(N),N为列表的长度
// 这是查找的最大次数

//二分查找

func BinarySearch(list[]int,search int)int{
	 low:=0
	 high:=len(list)-1
	 var mid int
	 var guess int
	 for low<=high{//只要范围没有缩小到只包含一个元素
	 	mid=(low+high)/2 //就检查中间的元素
	 	guess=list[mid]
	 	if guess==search{//找到了元素
	 		return mid
		}
	 	if guess>search{ //猜的数字大了
	 		high=mid-1
		}else{ //猜的数字小了
			low=mid+1
		}
	 }
	 return -1 //没有指定的元素
}

//大O表示法指出了算法有多快
//简单查找需要检查每个元素，因此需要执行n次操作。使用大O表示法，这个远行时间为O(n)
//大O表示法让你能够比较操作数，它指出了算法运行时间的增速。
//它表示的是操作数
//大O表示法说的是最糟的情形

//算法的速度指的并非时间，而是操作数的增速。
//谈论算法的速度时，我们说的是随着输入的增加，其运行时间将以什么样的速度增加。
//算法的运行时间用大O表示法表示
//O(log n)比O(n)快，当需要搜索的元素越多时，前者比后者快得越多