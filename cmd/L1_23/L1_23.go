package L1_23

func TwentyThree(words []string, i int) []string {
	return append(words[:i], words[i+1:]...)
}
