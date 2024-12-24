package main

import "fmt"

type StringBuilder struct {
	accum []string
}

func (sb *StringBuilder) Add(s string) {
	sb.accum = append(sb.accum, s)
}

func (sb *StringBuilder) Build() (res string) {
	for _, str := range sb.accum {
		res += str
	}

	return res
}

func main() {
	b := StringBuilder{}
	b.Add("A")
	b.Add("B")
	b.Add("C")
	fmt.Println(b.Build())
}