package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	z := 0
	i := len(matrix[0]) - 1
	for z < len(matrix) && i >= 0 {
		if matrix[z][i] > target {
			i--
		} else if matrix[z][i] < target { //else if不能换行
			z++
		} else {
			return true
		}

	}
	return false
}
func main() {
	var a = [][]int{}
	println(findNumberIn2DArray(a, 5))
}
