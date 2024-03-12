package day2

import "fmt"

func extractCharsFromString(str string, freqArr *[26]int, isCharExtractionDone chan bool) {
	for _, ch := range str {
		freqArr[ch-'a']++
	}
	isCharExtractionDone <- true
}

func Ex1() {
	strings := []string{"a", "quick", "brown", "fox", "jumped", "over", "a", "lazy", "dog"}
	chars := [26]int{}
	isCharExtractionDone := make(chan bool, 1000)

	for _, str := range strings {
		go extractCharsFromString(str, &chars, isCharExtractionDone)
		<-isCharExtractionDone
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", 'a'+i, chars[i])
	}
}
