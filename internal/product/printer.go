package product

import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	Name    string
	Brand   string
	Price   int
	InStock bool
}

const (
	template = `===== Alifshop =====
Товар:    {name}
Бренд:    {brand}
Цена:     {price} сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`
)

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", FormatPrice(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", FormatPrice(firstAmount))
	return text
}

func FormatPrice(price int) string {
	sum := price / 100
	tiin := price % 100

	s := strconv.Itoa(sum)
	n := len(s)

	var result string

	for i := 0; i < n; i++ {
		result += string(s[i])

		remains := n - i - 1

		if remains > 0 && remains%3 == 0 {
			result += " "
		}
	}

	if tiin == 0 {
		return result
	}
	return fmt.Sprintf("%s.%02d", result, tiin)
}
