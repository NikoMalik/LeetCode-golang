package main

func ispangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	var bits uint32 = 0

	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 'a' && sentence[i] <= 'z' {
			bits |= 1 << (sentence[i] - 'a')
		}
	}
	return bits == 0x3ffffff
}

func main() {
	t := ispangram("thequickbrownfoxjumpsoverthelazydog")

	println(t)

}

// explanation https://leetcode.com/problems/check-if-the-sentence-is-pangram/solutions/6366570/100-0ms-solution-with-explanation-by-nik-rnyr/
