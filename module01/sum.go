package module01

// Sum will sum up all of the numbers passed
// in and return the result
// func Sum(numbers []int) int {
// 	switch reflect.TypeOf(numbers).String() {
// 	case "[]int":
// 		var counter int
// 		for _, x := range numbers {
// 			counter += x
// 		}
// 		return counter
// 	default:
// 		return 0
// 	}
// }
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
