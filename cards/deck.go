package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // Deprecated
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	//Handle error
	// 	panic(err)
	// }
	defer file.Close()

	// Write Data to File
	_, err = file.Write([]byte(d.toString()))
	// if err != nil {
	// 	//Handle error
	// 	panic(err)
	// }

	// Optionally, you can use file.Sync() to flush writes to stable storage
	file.Sync()

	return err
}

func newDeckFromFile(filename string) deck {
	// bs, err := ioutil.ReadFile(filename) // Deprecated
	// if err != nil {
	// 	//Option #1 - log the error and return the call to newDeck()
	// 	//Option #2 - Log the error and entirely quit the program
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	//Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fileSize := fileInfo.Size()

	//Read data from the file
	bs := make([]byte, fileSize)
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
