package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val, found := cards[card]

	if !found {
		return 0
	}

	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	d := ParseCard(dealerCard) 

	switch {
	case shouldHit(c1, c2, d):
		return "H"
	case shouldSplit(c1, c2):
		return "P"
	case shouldStand(c1, c2, d):
		return "S"
	case shouldWin(c1, c2, d):
		return "W"
	}

	return ""
}

func shouldHit(c1, c2, d int) bool {
	switch {
	case c1+c2 <= 11,
		(c1+c2 >= 12 && c1+c2 <= 16) && d >= 7:
		return true
	}

	return false
}

func shouldStand(c1, c2, d int) bool {
	switch {
	case c1+c2 >= 17 && c1+c2 <= 20,
		c1+c2 >= 12 && c1+c2 <= 16,
		d == 11 || d == 10:
		return true
	}

	return false
}

func shouldSplit(c1, c2 int) bool {
	return c1 == 11 && c2 == 11
}

func shouldWin(c1, c2, d int) bool {
	return c1+c2 == 21 && !shouldStand(c1, c2, d)
}
