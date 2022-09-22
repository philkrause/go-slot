package main

type Data struct {
	Symbol []string
	Name   []string
	Payout []float32
}

// Reel Weight
var ReelOneWeight []int = []int{1, 1, 3, 3, 4, 5, 3}
var ReelTwoWeight []int = []int{1, 1, 3, 4, 4, 2, 5}
var ReelThreeWeight []int = []int{1, 1, 1, 4, 4, 3, 6}

// SymbolWeight ie. how often the symbol occurs
var aceWeight []int = []int{1, 1, 1}
var kingWeight []int = []int{1, 1, 1}
var queenWeight []int = []int{3, 3, 1}
var jackWeight []int = []int{3, 4, 4}
var spadeWeight []int = []int{4, 4, 4}
var diamondWeight []int = []int{5, 2, 3}
var jokerWeight []int = []int{3, 5, 6}

func Symbol() []string {

	data := Data{}
	data.Symbol = []string{"J", "Q", "K", "A", "H", "C", "D", "S", "B", "W"}

	return data.Symbol
}

func Names() []string {

	data := Data{}
	data.Name = []string{"Jacks", "Queens", "Kings", "Aces", "Hearts", "Clubs", "Diamond", "Spades", "Bonus", "Wilds"}

	return data.Name
}

func Payouts() []float32 {

	data := Data{}
	data.Payout = []float32{500, 100, 50, 20, 15, 10, 5, 2, 1}

	return data.Payout
}
