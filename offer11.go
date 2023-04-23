package main
//旋转数组中最小数字
func minArray(numbers []int) int {
	var min = 5001
	for i, v := range numbers {
		if v < min {
			min = numbers[i]
		}
	}
	//二分查找
	//var low=
	return min
}
