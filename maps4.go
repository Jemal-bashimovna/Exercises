package main

import (
	"fmt"
	"sort"
)

func main() {
	input := map[string]int{"Alice": 50, "Bob": 85, "Bobb": 85, "Charlie": -50, "Jane": 60}
	//result := RankParticipants(input)
	//fmt.Printf("%#v\n", result)

	test(input)
}
func RankParticipants(scores map[string]int) []string {

	participants := make([]string, 0, len(scores))

	for name := range scores {
		participants = append(participants, name)
	}
	sort.Slice(participants, func(i, j int) bool {
		if scores[participants[i]] > scores[participants[j]] {
			return true
		}
		if scores[participants[i]] == scores[participants[j]] && len(participants[i]) > len(participants[j]) {
			return true
		}
		return false
	})
	return participants

}

func test(sMap map[string]int) {
	fmt.Println(sMap)
	fmt.Println()
	p := make([]string, 0, len(sMap))
	for key := range sMap {
		p = append(p, key)
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p)-1-i; j++ {
			if sMap[p[j]] < sMap[p[j+1]] {
				p[j], p[j+1] = p[j+1], p[j]
			}
			if sMap[p[j]] == sMap[p[j+1]] && len(p[j]) < len(p[j+1]) {
				p[j], p[j+1] = p[j+1], p[j]
			}
		}
	}

	fmt.Println(p)
	fmt.Println(sMap)
}
