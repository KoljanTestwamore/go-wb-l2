package main

import "fmt"

// Можно и функции, но так можно более гибко настроить стратегии
type IStrategy interface {
	Run([]int) int
}

type SumStrategy struct {}
func (s SumStrategy) Run(vals []int) (res int) {
	for _, val := range vals {
		res += val
	}

	return res
}

type MultStrategy struct {}
func (s MultStrategy) Run(vals []int) (res int) {
	res = 1
	for _, val := range vals {
		res *= val
	}

	return res
}

type ArrayClient struct {
	Strategy IStrategy
	Data []int
}

func (a *ArrayClient) PerformCalculation() {
	fmt.Println(a.Strategy.Run(a.Data))
}

func main() {
	ac := ArrayClient{
		Data: []int{1, 2, 3, 4, 5},
		Strategy: MultStrategy{},
	}

	ac.PerformCalculation()
}