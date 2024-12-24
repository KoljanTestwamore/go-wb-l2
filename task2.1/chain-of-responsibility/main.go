package main

import "fmt"


type Suffix struct {
	next *Suffix
	sfx string
}

func NewSuffix(sfx string) Suffix {
	return Suffix{
		sfx: sfx,
	}
}

func (s1 *Suffix) SetNext(s2 *Suffix) {
	s1.next = (s2)
}

func (s *Suffix) GetWord(word string) (res string) {
	res = word
	sf := s

	for {
		res += sf.sfx
		sf = sf.next

		if (sf == nil) {
			return res
		}
	}
}

func main() {
	suffix1 := Suffix{
		sfx: "1",
	}

	suffix2 := Suffix{
		sfx: "2",
	}

	suffix2.SetNext(&suffix1)

	suffix3 := Suffix{
		sfx: "3",
	}

	suffix3.SetNext(&suffix2)

	fmt.Println(suffix3.GetWord("word"))

}