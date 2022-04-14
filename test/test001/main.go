package main

func main() {
	n := 3
	arr := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		arr[i+1] = true
	}
}

func perm(arr []int) []int {
	var tmp []int
	var newArr []int
	for i, x := range arr {
		tmp = []int{x}
		newArr = append(arr[0:i], arr[i+1:]...)
		tmp = append(tmp, perm(newArr)...)
	}
	return []int{}
}
