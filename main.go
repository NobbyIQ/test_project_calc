package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var count int
var resstr string

func Args(s []string) (int, int, error) {
	var isRoman bool
	var n1, n2 int
	var err error
	var lv, cv int

	for i := len(s[0]) - 1; i >= 0; i-- {
		cv = roman[string(s[0][i])]

		if cv != 0 {
			isRoman = true
			count = 1
		}

		if cv < lv {
			n1 -= cv
		} else {
			n1 += cv
		}
		lv = cv
	}
	lv = 0

	for i := len(s[2]) - 1; i >= 0; i-- {
		cv = roman[string(s[2][i])]

		if cv != 0 && count == 1 {
			isRoman = true
			count = 2
		} else if cv != 0 && count == 0 {
			isRoman = true
			count = 0
			return 0, 0, errors.New("Числа должны быть одного вида")
		}

		if cv < lv {
			n2 -= cv
		} else {
			n2 += cv
		}
		lv = cv
	}

	if !isRoman {
		n1, err = strconv.Atoi(s[0])
		if err != nil {
			return 0, 0, err
		}
	}
	if n1 > 10 || n1 < 0 {
		return 0, 0, errors.New("Числа должны быть от 1 до 10")
	}

	if !isRoman {
		n2, err = strconv.Atoi(s[2])
		if err != nil {
			return 0, 0, err
		}
	}
	if n2 > 10 || n2 < 0 {
		return 0, 0, errors.New("Числа должны быть от 1 до 10")
	}

	if count < 2 && count != 0 {
		count = 0
		return 0, 0, errors.New("Числа должны быть одного вида")
	}

	return n1, n2, nil
}

func Calc(s string) (int, error) {
	var res int

	c := strings.Split(s, " ")
	if len(c) < 2 {
		return res, errors.New("Недостаточно введенных аргументов")
	}

	if len(c) > 3 {
		return res, errors.New("Формат математической операции не удовлетворяет заданию")
	}

	n1, n2, err := Args(c)
	if err != nil {
		return res, err
	}

	switch c[1] {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		if n2 == 0 {
			return res, err
		} else {
			res = n1 / n2
		}
	}
	return res, nil
}

func main() {
	var expr string

	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			expr = scanner.Text()
			res, err := Calc(expr)
			if err != nil {
				fmt.Println(err)
			} else if count != 0 {
				if res <= 0 {
					fmt.Println(errors.New("В римских цифрах нет отрицательных значений"))
				}
				for res != 0 {
					if res-100 >= 0 {
						resstr += "C"
						res -= 100
					} else if res-90 >= 0 {
						resstr += "XC"
						res -= 90
					} else if res-50 >= 0 {
						resstr += "L"
						res -= 50
					} else if res-40 >= 0 {
						resstr += "XL"
						res -= 40
					} else if res-10 >= 0 {
						resstr += "X"
						res -= 10
					} else if res-9 >= 0 {
						resstr += "IX"
						res -= 9
					} else if res-5 >= 0 {
						resstr += "V"
						res -= 5
					} else if res-4 >= 0 {
						resstr += "IV"
						res -= 4
					} else if res-1 >= 0 {
						resstr += "I"
						res -= 1
					}
				}
				fmt.Println(resstr)
				count = 0
				resstr = ""
			} else {
				fmt.Println(res)
			}
		}
	}
}
