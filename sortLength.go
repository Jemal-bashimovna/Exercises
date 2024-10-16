package main

import (
	"fmt"
	"sort"
)

// Participant представляет участника с именем и очками
type Participant struct {
	Name  string
	Score int
}

// ByScoreAndName реализует интерфейс sort.Interface для сортировки участников
type ByScoreAndName []Participant

func (p ByScoreAndName) Len() int {
	return len(p)
}

func (p ByScoreAndName) Less(i, j int) bool {
	if p[i].Score == p[j].Score {
		return len(p[i].Name) > len(p[j].Name) // Сравниваем длину имен
	}
	return p[i].Score > p[j].Score // Сравниваем очки
}

func (p ByScoreAndName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// rankParticipants сортирует участников по очкам и длине имени
func rankParticipants(scores map[string]int) []string {
	participants := make([]Participant, 0, len(scores))

	// Заполняем срез Participant
	for name, score := range scores {
		participants = append(participants, Participant{Name: name, Score: score})
	}

	// Сортируем участников
	sort.Sort(ByScoreAndName(participants))

	// Извлекаем отсортированные имена
	rankedNames := make([]string, len(participants))
	for i, participant := range participants {
		rankedNames[i] = participant.Name
	}

	return rankedNames
}

func main() {
	// Пример использования
	scores := map[string]int{
		"Alice":   50,
		"Bob":     75,
		"Charlie": 75,
		"David":   60,
		"Eve":     50,
	}
	fmt.Println(scores)
	rankedParticipants := rankParticipants(scores)
	fmt.Println(rankedParticipants)
	fmt.Println(scores)
}
