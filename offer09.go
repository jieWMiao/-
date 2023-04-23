package main

func fib(n int) int {
	var a []int
	var c = 1000000007
	a = append(a, 0)
	a = append(a, 1)
	for i := 2; i <= 100; i++ {
		a = append(a, (a[i-1]+a[i-2])%c)
	}
	return a[n]

}
func main() {
	println(fib(0))
}
