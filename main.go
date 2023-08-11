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
	"I":     1,
	"II":    2,
	"III":   3,
	"IV":    4,
	"V":     5,
	"VI":    6,
	"VII":   7,
	"VIII":  8,
	"IX":    9,
	"X":     10,
	"XI":    11,
	"XII":   12,
	"XIII":  13,
	"XIV":   14,
	"XV":    15,
	"XVI":   16,
	"XVII":  17,
	"XVIII": 18,
	"XIX":   19,
}

var count int
var resstr string

func Args(s []string) (int, int, error) {
	var isRoman bool
	var n1, n2 int
	var err error

	for i, v := range roman {
		if s[0] == i {
			isRoman = true
			n1 = v
			count++
		}
	}

	for i, v := range roman {
		if s[2] == i {
			isRoman = true
			n2 = v
			count++
		}
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
				for i, v := range roman {
					if res == v {
						resstr = i
					}
				}
				fmt.Println(resstr)
				count = 0
			} else {
				fmt.Println(res)
			}
		}
	}
}
