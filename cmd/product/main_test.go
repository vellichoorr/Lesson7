package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []string
		shouldContain  bool
	}{
		{
			name: "Product with integer price and in stock",
			input: `iPhone 15 Pro
Apple
1500000
1`,
			expectedOutput: []string{
				"Товар:    iPhone 15 Pro",
				"Бренд:    Apple",
				"Цена:     1 500 000 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 125 000 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with decimal price and not in stock",
			input: `Samsung Galaxy S24
Samsung
1299999.99
0`,
			expectedOutput: []string{
				"Товар:    Samsung Galaxy S24",
				"Бренд:    Samsung",
				"Цена:     1 299 999.99 сум",
				"В наличии: false",
				"Рассрочка: 12 мес → 108 333.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with spaces in price",
			input: `MacBook Pro
Apple
25 000 000
1`,
			expectedOutput: []string{
				"Товар:    MacBook Pro",
				"Бренд:    Apple",
				"Цена:     25 000 000 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 2 083 333.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with small price with tiyin",
			input: `Ручка
Parker
2500.50
1`,
			expectedOutput: []string{
				"Товар:    Ручка",
				"Бренд:    Parker",
				"Цена:     2 500.50 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 208.37 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with zero price",
			input: `Test Product
Test Brand
0
0`,
			expectedOutput: []string{
				"Товар:    Test Product",
				"Бренд:    Test Brand",
				"Цена:     0 сум",
				"В наличии: false",
				"Рассрочка: 12 мес → 0 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with invalid price format",
			input: `Test Product
Test Brand
abc
1`,
			expectedOutput: []string{
				"вы вели не правильную сумму",
			},
			shouldContain: true,
		},
		{
			name: "Product with empty product name",
			input: `
Sony
999999
1`,
			expectedOutput: []string{
				"Товар:    ",
				"Бренд:    Sony",
				"Цена:     999 999 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 83 333.25 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with very large price",
			input: `Частный самолет
Gulfstream
999999999.99
0`,
			expectedOutput: []string{
				"Товар:    Частный самолет",
				"Бренд:    Gulfstream",
				"Цена:     999 999 999.99 сум",
				"В наличии: false",
				"Рассрочка: 12 мес → 83 333 333.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with one tiyin",
			input: `Тест
Test
1000.01
1`,
			expectedOutput: []string{
				"Товар:    Тест",
				"Бренд:    Test",
				"Цена:     1 000.01 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 83.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with invalid stock value defaults to false",
			input: `Product
Brand
1000
maybe`,
			expectedOutput: []string{
				"Товар:    Product",
				"Бренд:    Brand",
				"Цена:     1 000 сум",
				"В наличии: false",
				"Рассрочка: 12 мес → 83.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product exactly 12990000 sum",
			input: `ABC
a
12990000
1`,
			expectedOutput: []string{
				"Товар:    ABC",
				"Бренд:    a",
				"Цена:     12 990 000 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 1 082 500 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Product with 12990123.99 sum",
			input: `ABC
a
12990123.99
1`,
			expectedOutput: []string{
				"Товар:    ABC",
				"Бренд:    a",
				"Цена:     12 990 123.99 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 1 082 510.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Small price 5000 sum",
			input: `Ручка
Parker
5000
1`,
			expectedOutput: []string{
				"Товар:    Ручка",
				"Бренд:    Parker",
				"Цена:     5 000 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 416.66 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Exactly 1000000 sum",
			input: `Телефон
Samsung
1000000
0`,
			expectedOutput: []string{
				"Товар:    Телефон",
				"Бренд:    Samsung",
				"Цена:     1 000 000 сум",
				"В наличии: false",
				"Рассрочка: 12 мес → 83 333.33 сум/мес",
			},
			shouldContain: true,
		},
		{
			name: "Under thousand - 900 sum",
			input: `Ластик
Koh-i-Noor
900
1`,
			expectedOutput: []string{
				"Товар:    Ластик",
				"Бренд:    Koh-i-Noor",
				"Цена:     900 сум",
				"В наличии: true",
				"Рассрочка: 12 мес → 75 сум/мес",
			},
			shouldContain: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original stdin and stdout
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			}()

			// Create pipe for stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Write test input to pipe
			go func() {
				defer w.Close()
				w.Write([]byte(tt.input))
			}()

			// Create pipe for stdout
			outR, outW, _ := os.Pipe()
			os.Stdout = outW

			// Capture output
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, outR)
				outC <- buf.String()
			}()

			// Run main function
			main()

			// Close write side of stdout pipe
			outW.Close()

			// Get captured output
			output := <-outC

			// Check expected output
			for _, expected := range tt.expectedOutput {
				if tt.shouldContain && !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot output:\n%s", expected, output)
				}
			}
		})
	}
}

func TestMainWithVariousPriceFormats(t *testing.T) {
	tests := []struct {
		name          string
		priceInput    string
		shouldSucceed bool
	}{
		{"Integer price", "1000000", true},
		{"Decimal with dot", "1000000.50", true},
		{"Price with spaces", "1 000 000", true},
		{"Price with spaces and decimal", "1 000 000.99", true},
		{"Multiple spaces", "1  000  000", true},
		{"Negative price", "-1000", true},
		{"Invalid characters", "1000abc", false},
		{"Only letters", "million", false},
		{"Special characters", "$1000", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare full input
			input := "Test Product\nTest Brand\n" + tt.priceInput + "\n1"

			// Save original stdin and stdout
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			}()

			// Create pipe for stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Write test input
			go func() {
				defer w.Close()
				w.Write([]byte(input))
			}()

			// Create pipe for stdout
			outR, outW, _ := os.Pipe()
			os.Stdout = outW

			// Capture output
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, outR)
				outC <- buf.String()
			}()

			// Run main
			main()

			// Close stdout write side
			outW.Close()

			// Get output
			output := <-outC

			// Check if error message appears for invalid input
			hasError := strings.Contains(output, "вы вели не правильную сумму")
			if tt.shouldSucceed && hasError {
				t.Errorf("Expected successful parsing for %q, but got error", tt.priceInput)
			}
			if !tt.shouldSucceed && !hasError {
				t.Errorf("Expected error for invalid input %q, but parsing succeeded", tt.priceInput)
			}
		})
	}
}

func TestMainStockInputVariations(t *testing.T) {
	tests := []struct {
		name          string
		stockInput    string
		expectedStock string
	}{
		{"True as 1", "1", "true"},
		{"False as 0", "0", "false"},
		{"True string", "true", "true"},
		{"False string", "false", "false"},
		{"Invalid defaults to false", "maybe", "false"},
		{"Empty defaults to false", "", "false"},
		{"Negative number", "-1", "false"},
		{"Large number", "2", "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare input
			input := "Product\nBrand\n1000\n" + tt.stockInput

			// Save original stdin and stdout
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			}()

			// Setup stdin
			r, w, _ := os.Pipe()
			os.Stdin = r
			go func() {
				defer w.Close()
				w.Write([]byte(input))
			}()

			// Setup stdout
			outR, outW, _ := os.Pipe()
			os.Stdout = outW

			// Capture output
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, outR)
				outC <- buf.String()
			}()

			// Run main
			main()
			outW.Close()

			// Check output
			output := <-outC
			expectedOutput := "В наличии: " + tt.expectedStock
			if !strings.Contains(output, expectedOutput) {
				t.Errorf("For stock input %q, expected output to contain %q, got:\n%s",
					tt.stockInput, expectedOutput, output)
			}
		})
	}
}