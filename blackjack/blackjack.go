package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0 // for any other card value
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Check if the player has a pair of aces
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P" // Split
	}

	// Check if the player has a Blackjack
	switch {
	case (card1 == "ace" && card2 == "king") || (card1 == "king" && card2 == "ace"):
		// If dealer's card is not an ace, a figure, or a ten, automatically win
		switch {
		case dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king":
			return "W" // Automatically win
		}
		// Otherwise, stand
		return "S"
	}

	// Check if the player's total value is 11 or lower
	switch {
	case ParseCard(card1)+ParseCard(card2) <= 11:
		return "H" // Hit
	}

	// Check if the player's total value is within the range [12, 16]
	switch {
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		// If dealer's card is 7 or higher, hit
		switch {
		case ParseCard(dealerCard) >= 7:
			return "H" // Hit
		}
		return "S" // Stand
	}

	// Check if the player has a Blackjack and dealer's card is five
	switch {
	case ParseCard(card1)+ParseCard(card2) == 21 && dealerCard == "five":
		return "W" // Automatically win
	}

	// Otherwise, stand
	return "S"
}
