package product

import (
	"fmt"
	"testing"
)

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

func ExampleConverter_small_price() {
	var example = Product{
		Name:    "Ручка",
		Brand:   "Parker",
		Price:   5_000_00,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Ручка
	//Бренд:    Parker
	//Цена:     5 000 сум
	//В наличии: true
	//Рассрочка: 12 мес → 416.66 сум/мес
	//====================
}

func ExampleConverter_small_price_with_tiin() {
	var example = Product{
		Name:    "Карандаш",
		Brand:   "Faber-Castell",
		Price:   2_500_50,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Карандаш
	//Бренд:    Faber-Castell
	//Цена:     2 500.50 сум
	//В наличии: true
	//Рассрочка: 12 мес → 208.37 сум/мес
	//====================
}

func ExampleConverter_million_exact() {
	var example = Product{
		Name:    "Телефон",
		Brand:   "Samsung",
		Price:   1_000_000_00,
		InStock: false,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Телефон
	//Бренд:    Samsung
	//Цена:     1 000 000 сум
	//В наличии: false
	//Рассрочка: 12 мес → 83 333.33 сум/мес
	//====================
}

func ExampleConverter_under_thousand() {
	var example = Product{
		Name:    "Ластик",
		Brand:   "Koh-i-Noor",
		Price:   900_00,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Ластик
	//Бренд:    Koh-i-Noor
	//Цена:     900 сум
	//В наличии: true
	//Рассрочка: 12 мес → 75 сум/мес
	//====================
}

func ExampleConverter_one_tiyin() {
	var example = Product{
		Name:    "Тест",
		Brand:   "Test",
		Price:   1_000_01,
		InStock: true,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Тест
	//Бренд:    Test
	//Цена:     1 000.01 сум
	//В наличии: true
	//Рассрочка: 12 мес → 83.33 сум/мес
	//====================
}

func ExampleConverter_large_number() {
	var example = Product{
		Name:    "Машина",
		Brand:   "Tesla",
		Price:   999_999_999_99,
		InStock: false,
	}
	converter := Converter(example, Calculate(example.Price))
	fmt.Println(converter)
	//Output:
	//===== Alifshop =====
	//Товар:    Машина
	//Бренд:    Tesla
	//Цена:     999 999 999.99 сум
	//В наличии: false
	//Рассрочка: 12 мес → 83 333 333.33 сум/мес
	//====================
}

func TestConverter_edge_cases(t *testing.T) {
	tests := []struct {
		name     string
		product  Product
		contains []string
	}{
		{
			name: "Empty product name",
			product: Product{
				Name:    "",
				Brand:   "NoBrand",
				Price:   100_00,
				InStock: true,
			},
			contains: []string{"Товар:    \n", "Бренд:    NoBrand", "100"},
		},
		{
			name: "Very long product name",
			product: Product{
				Name:    "Очень длинное название товара которое может не поместиться",
				Brand:   "LongBrand",
				Price:   100_00,
				InStock: false,
			},
			contains: []string{"Очень длинное название товара которое может не поместиться", "false"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Converter(tt.product, Calculate(tt.product.Price))
			for _, substr := range tt.contains {
				if !contains(result, substr) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", substr, result)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
		 len(s) > 0 && len(substr) > 0 &&
		 findSubstring(s, substr) != -1)
}

func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
