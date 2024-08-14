package main

import (
	"fmt"
	"math/rand"
)

/*
Implement the following two-player card game:

Shuffle a standard deck (52 cards) and split it
in half so that each player starts with a pile of 26 cards.

Example deck of cards:
2,3,4,5,6,7,8,9,10,J,Q,K,A Club
2,3,4,5,6,7,8,9,10,J,Q,K,A Diamond
2,3,4,5,6,7,8,9,10,J,Q,K,A Heart
2,3,4,5,6,7,8,9,10,J,Q,K,A Spade

Every round, each player plays the top card from their pile.
The player with the higher card wins and adds both cards to
the bottom of their pile. We do not care about suits.
(Suits es el tipo, o sea corazones).

If at any point a player runs out of cards, the game is over
and that player loses.

If players tie (example: if both players play an 8 on the same turn),
then each player plays three cards. The players compete with
their third cards. The player with the higher competing card
wins and adds all of the cards that have been played during
the turn to the bottom of their pile. This process repeats as
necessary until the tie is resolved.

Tips:
1. You can use the compiler as necessary.
2. If you forget how to use a standard library function
in your coding language, please ask.
3. You can use a built-in shuffle function. If it doesn't exist
in your language, you can stub it out and try to implement it at the end.
*/

type Card struct {
	Name   string
	Type   string // Club, Diamond, Hear & Spade
	Weight int
}

type Player struct {
	Name       string
	Deck       []Card
	LatestCard Card
}

func createDeck() []Card {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	cardTypes := []string{"Club",
		"Diamond",
		"Hear",
		"Spade"}

	var deck []Card
	for _, cardType := range cardTypes {
		for j, card := range cards {
			deck = append(deck, Card{
				Name:   card,
				Type:   cardType,
				Weight: j + 1,
			})
		}
	}

	shuffle := make([]Card, len(deck))
	perm := rand.Perm(len(deck))
	for i, v := range perm {
		shuffle[v] = deck[i]
	}

	return shuffle
}

func (p *Player) pickCard(deck []Card) (bool, []Card) {
	l := len(deck)
	if l == 0 {
		return true, deck
	}
	p.LatestCard = deck[0]
	deck = deck[1:]
	return false, deck
}

func (p *Player) Pick3Cards(deck []Card) (bool, []Card) {
	l := len(deck)
	if l <= 3 {
		p.Deck = append(p.Deck, deck...)
		return true, deck
	}
	p.Deck = append(p.Deck, deck[0:3]...)
	deck = deck[3:]
	return false, deck
}

func main() {
	deck := createDeck()
	p1, p2 := &Player{Name: "Player1"}, &Player{Name: "Player2"}

	var end1, end2 bool
	for len(deck) != 0 {
		end1, deck = p1.pickCard(deck)
		end2, deck = p2.pickCard(deck)
		if end1 || end2 {
			if end1 == true {
				fmt.Println("the player1 loose because there are not more cards to pick")
			}
			fmt.Println("the player2 loose because there are not more cards to pick")
			break
		}

		if p1.LatestCard.Weight == p2.LatestCard.Weight {
			end1, deck = p1.Pick3Cards(deck)
			end2, deck = p2.Pick3Cards(deck)
			if end1 || end2 {
				if end1 == true {
					fmt.Println("the player1 loose because there are not more cards to pick")
				}
				fmt.Println("the player2 loose because there are not more cards to pick")
				break
			}
		}
		if p1.LatestCard.Weight > p2.LatestCard.Weight {
			p1.Deck = append(p1.Deck, p1.LatestCard, p2.LatestCard)
		} else {
			p2.Deck = append(p2.Deck, p2.LatestCard, p1.LatestCard)
		}
	}

	fmt.Println("**********")
	fmt.Println(deck)
	fmt.Println("**********")
	fmt.Println(*p1)
	fmt.Println("**********")
	fmt.Println(*p2)
}
