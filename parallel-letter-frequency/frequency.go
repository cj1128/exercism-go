package letter

type FreqMap map[rune]int

const testVersion = 1

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	chnl := make(chan FreqMap, len(texts))
	for _, s := range texts {
		go func(str string) {
			chnl <- Frequency(str)
		}(s)
	}

	//Create a master map to return
	returnMap := <-chnl

	for range texts[1:] {
		for char, cnt := range <-chnl {
			returnMap[char] += cnt
		}
	}
	return returnMap
}
