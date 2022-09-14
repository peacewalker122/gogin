package dummy

import "math/rand"

func Intgenerator(max int, min int) int {
	test := rand.Intn(int(max) - int(min)) + int(min)
	return test
}
func Namegenerator() string{
	charset := "abcdefghijklwnopqrstuvwxyz"
	c := charset[rand.Intn(len(charset))]
	return string(c)
}
func Currencygen() string{
	curset := []string{"IDR", "USD", "GBP", "YEN", "YUAN"}
	c := curset[rand.Intn(len(curset))]
	return string(c)
}