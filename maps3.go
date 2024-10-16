package main

import "fmt"

func main() {
	map1 := map[string]int{"banana": 10, "apple": 15}
	map2 := map[string]int{"banana": 7, "orange": 3, "apple": 15}

	result := MergeMaps(map1, map2)
	fmt.Printf("%#v\n", result)

}

func MergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)
	for val, i := range map1 {
		mergedMap[val] = i
	}
	for val, _ := range map2 {
		mergedMap[val] += map2[val]
	}
	return mergedMap
}
