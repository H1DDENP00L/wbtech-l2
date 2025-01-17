package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны.

Например:

"a4bc2d5e" => "aaaabccddddde"

"abcd" => "abcd"

"45" => "" (некорректная строка)

"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.

Например:

qwe\4\5 => qwe45 (*)

qwe\45 => qwe44444 (*)

qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку.

Написать unit-тесты.
*/

func unpackString(str string) (string, error) {
	runes := []rune(str)
	var result strings.Builder

	var numbers int
	var isEscape bool

	for i, v := range runes {
		if isEscape {
			result.WriteRune(v)
			isEscape = false
		} else if unicode.IsDigit(v) {
			numbers++
			factor, err := strconv.Atoi(string(v))
			if err != nil {
				return "", err
			}
			if i > 0 {
				result.WriteString(strings.Repeat(string(runes[i-1]), factor-1))
			} else {
				result.WriteRune(v)
			}
		} else if string(v) == `\` {
			isEscape = true
		} else {
			result.WriteRune(v)
		}
	}

	if str == "" {
		return "", nil
	} else if len(runes) == numbers {
		return "", fmt.Errorf("(некорректная строка)")
	} else {
		return result.String(), nil
	}

}

func main() {
	s, err := unpackString(`a4b1c`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
