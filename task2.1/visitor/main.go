package main

import "fmt"

type IBuilding interface {
	Accept(ISpy)
}

type ISpy interface {
	VisitHeadquarters(Headquarters) 
	VisitBase(Base) 
}

type Headquarters struct {
	title string
}
func (h *Headquarters) Accept(spy ISpy) {
	spy.VisitHeadquarters(*h)
}

type Base struct {
	size int
}
func (b *Base) Accept(spy ISpy) {
	spy.VisitBase(*b)
}

type JamesBond struct {}
func (jb *JamesBond) VisitHeadquarters(h Headquarters) {
	fmt.Println(h.title)
}

func (jb *JamesBond) VisitBase(b Base) {
	fmt.Println(b.size)
}

func main() {
	jb := JamesBond{}

	buildings := []IBuilding{&Headquarters{}, &Base{}}

	for _, building := range buildings {
		building.Accept(&jb)
	}
}
