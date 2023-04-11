package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, suit+" of "+val)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func getDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// func (d deck) shuffle() {
// 	for i := range d {
// 		newPosition := rand.Intn(len(d) - 1)
// 		/*
// 			In this code because the random package takes same seed value each time randomness soon becomes pattern hence to avoid it we use our own Random number implementation which is below
// 		*/

// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
