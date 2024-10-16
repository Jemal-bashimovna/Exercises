package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println(rand.Intn(100))
	//problem01()
	//problem02()
	//problem03()
	problem04()

}
func problem01() {
	cards := []string{"Zero", "One", "Two", "Three"}

	// for adding card
	//card := "Five"
	//cards = addCard(cards, card)
	//fmt.Println("New card successfully added")
	//
	//for i, val := range cards {
	//	fmt.Printf("Index [%d] => \"%s\"\n", i, val)
	//}

	// for removing card
	//index := 0
	//cards = removeCard(cards, index)
	//for i, val := range cards {
	//	fmt.Printf("Index [%d] => \"%s\"\n", i, val)
	//}

	// for shuffle deck
	cards = shuffleDeck(cards)
	fmt.Println(cards)

}
func addCard(deck []string, card string) []string {
	deck = append(deck, card)
	return deck
}

func removeCard(deck []string, index int) []string {
	if index >= len(deck) {
		fmt.Println("This index does not exist, select a number from 0 to ", len(deck)-1)
		return deck
	}
	deck = append(deck[:index], deck[index+1:]...)
	fmt.Println("The card successfully removed ")
	return deck
}

func shuffleDeck(deck []string) []string {

	newDeck := make([]string, len(deck))
	fmt.Printf("%T %#v\n", newDeck, newDeck)
	perm := rand.Perm(len(deck))
	fmt.Printf("%T %#v\n", perm, perm)

	//for i, randIndex := range perm {
	//	newDeck[i] = deck[randIndex]
	//	fmt.Println(newDeck[i])
	//}

	for i := 0; i < len(deck); i++ {
		newDeck[i] = deck[perm[i]]
		fmt.Println(newDeck[i])
	}
	return newDeck
}

func problem02() {

	temps := []float64{23, 20, 36.4, 35.2, 29.6, 40}
	minTemp, maxTemp, avTemp := analyzeTemperatures(temps)

	fmt.Printf("Minimal tepmerature: %f\nMaximal temperature: %f\nAvarage temperature: %f\n", minTemp, maxTemp, avTemp)
}

func analyzeTemperatures(temps []float64) (float64, float64, float64) {
	minTemp := temps[0]
	maxTemp := temps[0]
	var avTemp float64 = 0

	for _, value := range temps {
		if minTemp >= value {
			minTemp = value
		}
		if maxTemp <= value {
			maxTemp = value
		}
		avTemp = avTemp + value
	}
	avTemp = avTemp / float64(len(temps))

	return minTemp, maxTemp, avTemp
}

func problem03() {
	person := []string{"Anna", "John", "Mark"}
	newPerson := "Peter"
	person = addParticipant(person, newPerson)
	fmt.Println("New participant successfully added\n", person)

	pIndex := 0
	person = removeParticipant(person, pIndex)
	for i, p := range person {
		fmt.Printf("Index [%d] => \"%s\"\n", i, p)
	}

	name := "lk"
	result := findParticipant(person, name)
	fmt.Println("Index of participant = ", result)
}

func addParticipant(participants []string, participant string) []string {
	participants = append(participants, participant)
	return participants
}

func removeParticipant(participants []string, index int) []string {
	if index >= len(participants) {
		fmt.Println("This index does not exist, select a number from 0 to ", len(participants)-1)
		return participants
	}
	participants = append(participants[:index], participants[index+1:]...)
	fmt.Println("New Participant removed successfully")
	return participants

}

func findParticipant(participants []string, name string) int {
	for i, p := range participants {
		if name == p {
			fmt.Printf("%s => %s\n", name, p)
			return i
		}
	}
	return -1
}

func problem04() {
	numSlice := []int{1, 1, 2, 2, 3, 4, 5, 1, 2}

	//target := 1
	//numSlice = removeAll(numSlice, target)
	//fmt.Println(numSlice)

	numSlice = compress(numSlice)
	fmt.Println(numSlice)

	//numSlice = removeDuplicates(numSlice)
	//fmt.Println(numSlice)

}

func removeAll(nums []int, target int) []int {
	var newNumbers []int

	for i := 0; i < len(nums); i++ {
		if target != nums[i] {
			newNumbers = append(newNumbers, nums[i])
		}

	}
	return newNumbers
}

func compress(nums []int) []int {

	newNums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			newNums = append(newNums, nums[i])
		}
	}
	return newNums
}

func removeDuplicates(nums []int) []int {
	var newNums []int

	for i := 0; i < len(nums); i++ {
		exist := false
		for j := 0; j < len(newNums); j++ {
			if nums[i] == newNums[j] {
				exist = true
				break
			}
		}
		if !exist {
			newNums = append(newNums, nums[i])
		}
	}
	return newNums
}
