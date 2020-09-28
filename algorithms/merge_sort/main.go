package main

import "fmt"

func sortFunc(slice []int) []int {
	var index int
	newSorted := []int{}
	if len(slice) < 1 {
		return newSorted
	}
	small := slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] < small {
			small = slice[i]
			index = i
		}
	}
	newSorted = append(newSorted, small)
	slice = remove(slice, index)
	fmt.Println(newSorted)
	return sortFunc(slice)
}

func remove(ints []int, i int) []int {
	return append(ints[:i], ints[i+1:]...)
}
func main() {
	fmt.Println(sortFunc([]int{3, 4, 2, 1}))
}

// package main

// import "fmt"

// func main() {
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}
// }
