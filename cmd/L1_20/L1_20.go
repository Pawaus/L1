package L1_20

import (
	"fmt"
	"strings"
)

func reverseSlice(str_slice []string) (rev_slc []string) {
	for i := len(str_slice) - 1; i >= 0; i-- {
		// go from the end of the slice and append, that's why we are getting reverse slice
		rev_slc = append(rev_slc, str_slice[i])
	}
	return
}

func Twenty() {
	str := "snow dog sun"

	words := strings.Split(str, " ")

	words = reverseSlice(words)

	var reversed_words_str string
	for _, val := range words {
		reversed_words_str += val + " "
	}

	reversed_words_str = reversed_words_str[:len(reversed_words_str)-1]

	fmt.Println(reversed_words_str)
}
