package main

import (
	"strconv"
	"strings"
)

func isLetter(r rune) bool {
	return r >= rune('a') && r <= rune('z')
}

func isSlash(r rune) bool {
	return r == '\\'
}

func Unpack(str string) (res string) {
	// сделаем простой автомат
	state := 0
	prev := ""
	count := 0

	for _, r := range str {
		switch state {
		// начальное состояние
		case 0:
			if (isSlash(r)) {
				state = 3
				continue
			}
			if isLetter(r) {
				prev = string(r)
				state = 1

				continue
			}
			// fmt.Println("Bad string!")
			return ""
		// только что была буква
		case 1:
			if (isSlash(r)) {
				state = 3
				continue
			}
			if (isLetter(r)) {
				res += prev
				prev = string(r)

				continue
			}

			i, err := strconv.Atoi(string(r))
			if err != nil {
				// fmt.Println("Error parsing number!")
				return ""
			}
			count = i
			state = 2
		// только что была цифра
		case 2:
			if (isSlash(r)) {
				state = 3
				continue
			}
			if (isLetter(r)) {
				res += strings.Repeat(prev, count)
				prev = string(r)

				state = 1
				continue
			}

			i, err := strconv.Atoi(string(r))
			if err != nil {
				// fmt.Println("Error parsing number!")
				return ""
			}
			count = count * 10 + i
		// только что был слеш
		case 3:
			res += prev
			prev = string(r)
			state = 1

		default:
			// fmt.Println("Bad string")
			return ""
		}
	}

	switch state {
	case 0:
		return ""
	case 1:
		res += prev
	case 2:
		res += strings.Repeat(prev, count)
	default:
		// fmt.Println("Bad string")
		return ""
	}

	return res
}

func main() {
	Unpack("a4bc2d5e")
	Unpack("abcd")
	Unpack("45")
	Unpack("qwe\\4\\5")
	Unpack("qwe\\45")
	Unpack("qwe\\\\5")
}