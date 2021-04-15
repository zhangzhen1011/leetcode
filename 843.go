import "math/rand"

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(wordlist []string, master *Master) {
	for i := 0; i < 10; i++ {
		s := wordlist[rand.Intn(len(wordlist))]
		k := master.Guess(s)
		if k == 6 {
			return
		}

		tmp := make([]string, 0)
		for _, val := range wordlist {
			if similarity(val, s) == k {
				if s == val {
					continue
				}
				tmp = append(tmp, val)
			}
		}
		wordlist = tmp
	}
}

func similarity(s1, s2 string) int {
	c := 0
	for i := 0; i < 6; i++ {
		if s1[i] == s2[i] {
			c++
		}
	}
	return c
}
