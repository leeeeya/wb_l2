// поиск анаграм по словарю

package main

import (
	"fmt"
	"sort"
	"strings"
)

func searchAnagrams(dict []string) map[string][]string {
	// мапа для хранения анаграм, где ключ - первое встретившееся в словаре слово из множества, значение - слайс,
	// каждый элемент которого - анаграма к ключу
	anagrams := make(map[string][]string)

	// временная мапа для поиска анаграм, где ключ - строка из отсортированных символов, значение - первое встретившееся в словаре слово из множества,
	// состоящее из символов ключа
	sorted := make(map[string]string)

	for _, word := range dict {
		// приведение слова к нижнему регистру
		word = strings.ToLower(word)
		// приведение слова к слайсу строк, где каждый элемент - символ строки
		slWord := strings.Split(word, "")
		// сортировка символов в слайсе
		sort.Slice(slWord, func(i, j int) bool {
			return slWord[i] < slWord[j]
		})
		// приведение сортированного слайса обратно к строке
		sortedWord := strings.Join(slWord, "")
		// если слово - первое из множества анаграм, оно вносится во временную мапу и будет ключом в мапе анаграм
		if _, ok := sorted[sortedWord]; !ok {
			sorted[sortedWord] = word
		} else if ok {
			// если такое слово уже встречалось, то оно вносится в спивок анаграм
			anagrams[sorted[sortedWord]] = append(anagrams[sorted[sortedWord]], word)
		}
	}
	// сортировка слайса анаграм в мапе
	for _, v := range anagrams {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	return anagrams
}

func main() {

	dict := []string{"Материк", "Апельсин", "Рост", "Спаниель", "Сорт", "Телефон",
		"Листок", "Торс", "Столик", "Трос", "Метрика", "Слиток"}

	anagrams := searchAnagrams(dict)

	fmt.Println(anagrams)

}
