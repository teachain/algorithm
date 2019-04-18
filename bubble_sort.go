package algorithm

func BubbleSort(src []int) {
	length := len(src)
	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if src[j] > src[j+1] {
				temp := src[j+1]
				src[j+1] = src[j]
				src[j] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
