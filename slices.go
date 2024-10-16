package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//problem1()
	problem2()

}

func problem1() {

	cards := []string{"First", "Second", "Third", "Fourth"}
	card := "Fifth"
	addCard01(cards, card)
	index := 1
	removeCard01(cards, index)

	shuffleDeck01(cards)

}

func addCard01(deck []string, card string) {
	deck = append(deck, card)
	fmt.Printf("Type %T value %#v\n", deck, deck)
}

func removeCard01(deck []string, index int) {
	deck = append(deck[:index], deck[index+1:]...)
	fmt.Printf("Type %T values %#v\n", deck, deck)
}

func shuffleDeck01(deck []string) {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	//fmt.Printf("%T %#v\n", rand, rand)
	// fmt.Println(rand.Intn(10))
	newSlice := make([]string, len(deck))
	perm := rand.Perm(len(deck)) // rand.Perm()возвращает случайную перестановку чисел от 0 до N (не включая N)
	//fmt.Printf("type %T value %#v\n", perm, perm)
	for i, randI := range perm {
		newSlice[i] = deck[randI]
	}
	fmt.Println(newSlice)
	//fmt.Printf("%T %#v", newSlice, newSlice)
}

func problem2() {
	temperatures := []float64{30.8, 20, 23.4, 34, 34.7, 28.6, 36}
	fmt.Println(temperatures)
	minTemp, maxTemp, avTemp := analyzeTemperatures01(temperatures)
	fmt.Printf("min temp: %f, max temp: %f, o temp: %f", minTemp, maxTemp, avTemp)
}
func analyzeTemperatures01(temps []float64) (float64, float64, float64) {
	sort.Float64s(temps)
	minTemp := temps[0]
	maxTemp := temps[len(temps)-1]
	var avTemp float64 = 0
	for _, val := range temps {
		avTemp += val
	}

	return minTemp, maxTemp, avTemp / float64(len(temps))
}
