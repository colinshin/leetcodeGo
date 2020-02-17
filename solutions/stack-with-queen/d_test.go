package stack_with_queen

import "fmt"

func Example() {
	obj := Constructor()
	obj.Push(8)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)

	// output:
	// 8
	// -1
	// true
}
