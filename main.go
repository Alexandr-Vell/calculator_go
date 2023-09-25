package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputValues() (string, string, string) {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, " ")

	if len(parts) < 3 {
		fmt.Println("Ошибка: Вводимые данные не удовлетворяют условию, строка не является математической операцией")
		os.Exit(0)
	}

	if len(parts) > 3 {
		fmt.Println("Ошибка: Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		os.Exit(0)
	}

	rome, arabic := RomOrArabic(parts[0], parts[2])

	if len(parts) == 3 {
		if arabic || rome {
			return parts[0], parts[1], parts[2]
		} else {
			fmt.Println("Ошибка: Обе цифры должны быть либо арабские, либо римские одновременно")
			os.Exit(0)
			return "0", "+", "0"
		}
	} else {
		fmt.Println("Ошибка: Вводимые данные не удовлетворяют условию")
		os.Exit(0)
		return "0", "+", "0"
	}

}

func RomOrArabic(a, b string) (bool, bool) {
	arrayRom := generNum(100)

	/* generNum(N) генерирует от 1 до N римские цифры (100 выбрано рандомно, но > 10, чтобы "римские" ошибки
	по диапазону от 1 до 10 имели возможность сработать) */

	var valueA bool
	var valueB bool
	var valueAB bool
	var equalityAB bool

	for _, element := range arrayRom {
		if element == a {
			valueA = true
			break
		} else {
			valueA = false
		}

	}

	for _, element := range arrayRom {
		if element == b {
			valueB = true
			break
		} else {
			valueB = false
		}

	}

	//возвращает true, если оба значения римские цифры
	if valueA && valueB {
		valueAB = true
	} else {
		valueAB = false
	}

	//возвращает true, если оба значения арабские цифры
	if (valueA == false && valueB == true) || (valueA == true && valueB == false) || valueAB == true {
		equalityAB = false
	} else {
		equalityAB = true
	}

	return valueAB, equalityAB
}

func transformInt(a, operator, b string) (int, string, int) {

	numA, err := strconv.Atoi(a)
	if err != nil || numA < 1 || numA > 10 {
		fmt.Println("Первое значение должно быть целым числом и лежать в диапазоне от 1 до 10 включительно")
		os.Exit(0)
		return 0, "+", 0
	}

	numB, err := strconv.Atoi(b)
	if err != nil || numB < 1 || numB > 10 {
		fmt.Println("Второе значение должно быть целым числом и лежать в диапазоне от 1 до 10 включительно")
		os.Exit(0)
		return 0, "+", 0

	}

	return numA, operator, numB

}

func transformRom(a, operator, b string) (int, string, int) {

	rome, arabic := RomOrArabic(a, b)

	var numA int
	var numB int

	if rome {
		numA = fromRomInArabic(a)

		if numA < 1 || numA > 10 {
			fmt.Println("Первое значение должно быть целым числом и лежать в диапазоне от Ⅰ до Ⅹ включительно")
			os.Exit(0)
			return 0, "+", 0

		}

		numB = fromRomInArabic(b)

		if numB < 1 || numB > 10 {
			fmt.Println("Второе значение должно быть целым числом и лежать в диапазоне от Ⅰ до Ⅹ включительно")
			os.Exit(0)
			return 0, "+", 0
		}

	}

	if arabic {
		numA, operator, numB = transformInt(a, operator, b)
	}

	return numA, operator, numB

}

func calculate(a, operator, b string) int {

	numA, operator, numB := transformRom(a, operator, b)

	switch operator {
	case "+":
		return numA + numB
	case "-":
		return numA - numB
	case "*":
		return numA * numB
	case "/":
		return numA / numB //не добавляю проверку на numB != 0 потому что transformInt просто не пропустит сюда ноль
	default:
		fmt.Println("Ошибка: Такого оператора я еще не знаю :(")
		os.Exit(0)
		return 0
	}
}

func responseLang(a, operator, b string) string {

	total := calculate(a, operator, b)
	rome, arabic := RomOrArabic(a, b) // если rome - true, то это римские

	if rome && total > 0 {
		return fromArabInRome(total)
	} else if arabic {
		return strconv.Itoa(total)
	} else {
		fmt.Println("Ошибка: Ответ меньше или равен нулю, в римских цифрах отсутствует ноль и отрицательные числа")
		os.Exit(0)
		return ""
	}
}

func fromRomInArabic(a string) int {

	mapFullRome := map[string]int{
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	arrayRome := strings.Split(a, "")

	result := mapFullRome[string(arrayRome[len(a)-1])] // сразу последнее кидаем в result

	for i := len(a) - 2; i >= 0; i-- {
		arabic := mapFullRome[string(arrayRome[i])] // предпоследнее

		if arabic < mapFullRome[string(arrayRome[i+1])] {
			result -= arabic
		} else {
			result += arabic
		}

	}

	return result

}

func fromArabInRome(a int) string {
	arabicValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanValues := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(arabicValues); i++ {
		for a >= arabicValues[i] {
			result += romanValues[i]
			a -= arabicValues[i]
		}
	}

	return result
}

func generNum(a int) []string {
	arrayValues := []string{}

	for i := 0; i <= a; i++ {
		arrayValues = append(arrayValues, fromArabInRome(i))
	}

	return arrayValues

}

func main() {
	fmt.Println("Введите 2 значения и оператор (Например: 1+1): ")
	a, operator, b := inputValues()
	fmt.Println(responseLang(a, operator, b))

}
