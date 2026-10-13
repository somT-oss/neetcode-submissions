import "maps"

func isAnagram(s string, t string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, val := range s {
		_, ok := map1[val]
		if !(ok){
			map1[val] = 1
		} else {
			map1[val] += 1
		}
	}

	for _, val := range t {
		_, ok := map2[val]
		if !(ok){
			map2[val] = 1
		} else {
			map2[val] += 1
		}
	}

	return maps.Equal(map1, map2)
}
