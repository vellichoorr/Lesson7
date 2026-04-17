package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"product/internal/product"
)

func main() {
	var (
		productInfo product.Product
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Product info: ")
	productInfo.Name, _ = reader.ReadString('\n')

	productInfo.Name = strings.TrimSuffix(productInfo.Name, "\n")

	fmt.Print("Brand: ")
	productInfo.Brand, _ = reader.ReadString('\n')
	productInfo.Brand = strings.TrimSuffix(productInfo.Brand, "\n")
	fmt.Print("Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSuffix(priceStr, "\n")
	priceStr = strings.ReplaceAll(priceStr, " ", "")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("вы вели не правильную сумму")
	}

	fmt.Print("In stock? (0-false,1-true): ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSuffix(stockStr, "\n")
	productInfo.InStock, err = strconv.ParseBool(stockStr)
	if err != nil {
		fmt.Println(err)
	}

	productInfo.Price = int(price * tiinToSum)

	calculatedAmount := product.Calculate(productInfo.Price)

	converted := product.Converter(productInfo, calculatedAmount)
	fmt.Println(converted)
}

const (
	tiinToSum = 100
)

//10 000 000.99
//10 000 000 99
