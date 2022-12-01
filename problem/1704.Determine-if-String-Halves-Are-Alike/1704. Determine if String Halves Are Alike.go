package _704_Determine_if_String_Halves_Are_Alike

func halvesAreAlike(s string) bool {
	i := 0
	j := len(s) - 1
	a := 0
	b := 0

	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for i < j {
		if vowels[s[i]] {
			a++
		}
		if vowels[s[j]] {
			b++
		}
		i++
		j--
	}

	return a == b
}
