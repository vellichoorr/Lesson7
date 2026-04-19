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

// formatMoney форматирует сумму в тийинах в читаемый вид
func formatMoney(tiyin int) string {
	// Переводим тийины в сумы
	sum := tiyin / 100
	tiynRemainder := tiyin % 100

	// Преобразуем целую часть в строку
	sumStr := strconv.Itoa(sum)

	// Добавляем разделители тысяч
	formattedSum := addThousandSeparators(sumStr)

	// Добавляем тийины только если они есть
	if tiynRemainder > 0 {
		return fmt.Sprintf("%s.%02d", formattedSum, tiynRemainder)
	}
	return formattedSum
}

// addThousandSeparators добавляет пробелы между разрядами
func addThousandSeparators(s string) string {
	if len(s) <= 3 {
		return s
	}

	var result []byte
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result = append(result, ' ')
		}
		result = append(result, byte(ch))
	}
	return string(result)
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatMoney(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatMoney(firstAmount))
	return text
}
