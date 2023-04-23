package main

import (
	"flag"
)

var j = flag.Int("j", 0, "")

// offer3找重复数字
func findRepeatNumber(nums []int) int {
	var site = make(map[int]int) //第一行代码并没有对map进行一个初始化，而却对其进行写入操作，就是对空指针的引用，这将会造成一个painc。所以，得记得用make函数对其进行分配内存和map是引用类型
	for _, v := range nums {
		site[v] = 0
	}
	for _, v := range nums {

		site[v]++
		if site[v] >= 2 {
			return v
		}
	}
	return 0
}
func main() {
	b := []int{3, 4, 2, 1, 1, 0} // []int{2, 3, 1, 0, 2, 5, 3}方法中申明的形参是未指定数组大校的所以只能传这个
	var z = findRepeatNumber(b)
	println(z)
}
