package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func isRoman(s string) (ok bool) {
	_, ok = romanToArabic[s]
	return
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите пример: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	var (
		romflag            bool
		num1, num2, result int
	)

	if len(parts) != 3 {
		panic("Не подходящий формат ввода")
	}

	a, oper, b := parts[0], parts[1], parts[2]

	if (isRoman(a) && !isRoman(b)) || (!isRoman(a) && isRoman(b)) {
		panic("Нельзя вводить арабские и римские числа одновременно")
	}

	if isRoman(a) {
		num1, num2, romflag = romanToArabic[a], romanToArabic[b], true
	} else {
		num1, _ = strconv.Atoi(a)
		num2, _ = strconv.Atoi(b)
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Числа не соответсвуют диапазону от 1 до 10")
	}

	switch oper {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic("Данной операции не существует")
	}

	if romflag {
		if result <= 0 {
			panic("После выполнения операции с римскими цифрами, ответ должен быть положительным")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Printf("Результат: %v", result)
	}
}
