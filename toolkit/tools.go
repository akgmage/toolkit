package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ0123456789_+"

type Tools struct {
}
// Generate random string of length n
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x % y]
	}
	return string(s)
}