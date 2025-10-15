package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch(card) {
    case "two": 
    	return	2
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
    case "ace": 
    	return 11
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerVal := ParseCard(dealerCard)
                                              
	switch {
    case card1 == "ace" && card2 == "ace":
        return "P"
    case playerSum == 21:
        if dealerVal == 10 || dealerVal == 11 {
        	return "S"
        }
        return "W"
    case 17 <= playerSum && playerSum <= 20:
        return "S"
    case 12 <= playerSum && playerSum <= 16:
        if dealerVal >= 7 {
            return "H"
        }
        return "S"
    default:
        return "H"
    }
}
