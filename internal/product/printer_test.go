package product

import "fmt"

func ExampleConverter() {
	var example = Product{
		Name:    "ABC",
		Brand:   "a",
		Price:   12_990_000_00,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    ABC
	//Бренд:    a
	//Цена:     12 990 000 сум
	//В наличии: true
	//Рассрочка: 12 мес → 1 082 500 сум/мес
	//====================
}
func ExampleConverter_tiin() {
	var example = Product{
		Name:    "ABC",
		Brand:   "a",
		Price:   12_990_123_99,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    ABC
	//Бренд:    a
	//Цена:     12 990 123.99 сум
	//В наличии: true
	//Рассрочка: 12 мес → 1 082 510.33 сум/мес
	//====================
}
