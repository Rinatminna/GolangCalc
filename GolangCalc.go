package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input:")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	answer, err := calc(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:\n" + answer)
}

func calc(input string) (string, error) {
	var isRoman bool
	var result int

	inputSplit := strings.Fields(input)
	if len(inputSplit) != 3 {
		return "", errors.New("Неверно введённое выражение")
	}

	firstNumber, err := strconv.Atoi(inputSplit[0])
	secondNumber, err2 := strconv.Atoi(inputSplit[2])

	if err != nil || err2 != nil {
		firstNumber, err = romanToArab(inputSplit[0])
		if err != nil {
			return "", errors.New("Ошибка считывания цифр")
		}
		secondNumber, err = romanToArab(inputSplit[2])
		if err != nil {
			return "", errors.New("Ошибка считывания цифр")
		}
		isRoman = true
	}

	if firstNumber < 1 || firstNumber > 10 || secondNumber < 1 || secondNumber > 10 {
		return "", errors.New("Введенные числа выходят из диапазона ожидаемых")
	}

	sign := inputSplit[1]
	switch sign {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		result = firstNumber / secondNumber
	default:
		return "", errors.New("Неверно введен арифметический знак")
	}

	if isRoman {
		if result < 1 {
			return "", errors.New("Результат меньше 1, невозможно конвертировать в римские")
		}
		return arabToRome(result), nil
	}
	return strconv.Itoa(result), nil
}

func romanToArab(romanInput string) (int, error) {
	arab := []int{10, 9, 5, 4, 1}
	roman := []string{"X", "IX", "V", "IV", "I"}
	result := 0

	for i := 0; i < len(arab); i++ {
		for strings.HasPrefix(romanInput, roman[i]) {
			result += arab[i]
			romanInput = romanInput[len(roman[i]):]
		}
	}
	if len(romanInput) > 0 {
		return 0, errors.New("Неверный римский символ")
	}
	return result, nil
}

func arabToRome(arabInput int) string {
	arab := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; i < len(arab); i++ {
		value := arabInput / arab[i]
		for j := 0; j < value; j++ {
			result += roman[i]
		}
		arabInput %= arab[i]
	}
	return result
}
