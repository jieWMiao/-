package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var n = []int{}
	var p *ListNode = head
	for p != nil {
		n = append(n, p.Val)
		p = p.Next
	}
	i := 0
	j := len(n) - 1
	for i < j { //两个指针置换数组
		t := n[i]
		n[i] = n[j]
		n[j] = t
		i++
		j--
	}

	return n
}
func main() {

}
