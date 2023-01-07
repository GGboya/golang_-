func rev(num []int) []int {
	// 翻转切片的函数，便于自己使用
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
	return num
}

func sum(nums []int) (s int) {
	// 对nums数组进行求和
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return
}