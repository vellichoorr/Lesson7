package product

import "fmt"

func ExampleCalculate() {
	calculate := Calculate(1_000_00)
	fmt.Println(calculate)
	//Output:
	//8333
}

func ExampleCalculate_negative() {
	calculate := Calculate(-1_000_00)
	fmt.Println(calculate)
	//Output:
	//-8333
}

func ExampleCalculate_zero() {
	calculate := Calculate(0)
	fmt.Println(calculate)
	//Output:
	//0
}
