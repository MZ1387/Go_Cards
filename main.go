package main

func main() {
	cards := deck{newCard(), "Ace of Base"}
	cards = append(cards, "King of Hearts")
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
