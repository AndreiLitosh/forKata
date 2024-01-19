package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator(text string) []string {
	elements := []string{}
	text = strings.TrimSpace(text)
	elementsArr := strings.Split(text, " ")
	for _, el := range elementsArr {
		if el != "\n" {
			elements = append(elements, el)
		}
	}
	if len(elements) != 3 {
		panic("неверный формат ввода")
	}
	return elements
}

func romanOrInt(a string) (string, int) {
	rom := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	switch a {
	case "+":
		return a, 0
	case "-":
		return a, 0
	case "*":
		return a, 0
	case "/":
		return a, 0
	}
	for i, v := range rom {
		if i == a {
			el := v
			return a, el
		}
	}
	element1, err := strconv.Atoi(a)
	if err != nil || element1 < 1 || element1 > 10 {
		panic("Неверное число 1")
	}
	return " ", element1

}

func operation(el1s, sumbol, el2s string, el1, el2 int) string {
	var result int
	if el1s == " " && el2s != " " {
		panic("разные системы счисления")
	}
	if el1s != " " && el2s == " " {
		panic("разные системы счисления")
	}
	if el1s == " " {
		switch sumbol {
		case "+":
			result = el1 + el2
			return strconv.Itoa(result)
		case "-":
			result = el1 - el2
			return strconv.Itoa(result)
		case "*":
			result = el1 * el2
			return strconv.Itoa(result)
		case "/":
			result = el1 / el2
			return strconv.Itoa(result)
		default:
			panic("Не верный знак")
		}
	}
	switch sumbol {
	case "+":
		result = el1 + el2
		return intToRom(result)
	case "-":
		result = el1 - el2
		if result < 0 {
			panic("Отрицательное Римкое число")
		}
		return intToRom(result)
	case "*":
		result = el1 * el2
		return intToRom(result)
	case "/":
		result = el1 / el2
		return intToRom(result)
	default:
		panic("Не верный знак")
	}
}

func intToRom(result int) string {
	var rim string
	var numb = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var rimsl = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	index := len(rimsl) - 1
	for result > 0 {
		for numb[index] <= result {
			rim += rimsl[index]
			result -= numb[index]
		}
		index -= 1
	}
	return rim
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите число")
	text, _ := reader.ReadString('\n')
	a := calculator(text)
	s1, i1 := romanOrInt(a[0])
	s2, _ := romanOrInt(a[1])
	s3, i3 := romanOrInt(a[2])
	b := operation(s1, s2, s3, i1, i3)
	fmt.Println(b)
}
