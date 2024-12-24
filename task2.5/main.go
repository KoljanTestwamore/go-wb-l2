package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagramms(words *[]string) (res map[string][]string) {
	indexLabel := map[string]string{}
	res = make(map[string][]string)

	for _, w := range *words {
		word := strings.ToLower(w)

		indexMap := map[rune]int{}

		for _, char := range word {
			indexMap[char]++
		}

		// Кодируем индекс. Sprint кодирует мапы в отсортированном порядке
		// А так, конечно, можно было собрать код вручную
		index := fmt.Sprint(indexMap)

		value, ok := indexLabel[index]
		if ok {
			res[value] = append(res[value], word)
			continue
		}
		indexLabel[index] = word
		res[word] = []string{word}
	}

	for key, slice := range res {
		if len(slice) == 1 {
			delete(res, key)
			continue
		}
		sort.Strings(slice)
	}

	return res
}

func main() {
	// 4. Массив должен быть отсортирован по возрастанию
	// - Непонятно, что имелось в виду. Массив сам сортировать
	// не имеет смысла, отсортировал вывод
	fmt.Print(FindAnagramms(&[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "приколюшка"}))
}