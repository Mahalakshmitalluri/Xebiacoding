package main

import (
	"fmt"
	"strings"
)

var counter int
var myrune []rune

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c rune) bool {
	return ('0' <= c && c <= '9')
}

func singleRune(temp []rune) rune {
	return temp[0]
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func NewSkipString(val int, input string) Strconverter {

	return Strconverter{Value: val, Str: input}

}

type Strconverter struct {
	Value int
	Str   string
}

func (s Strconverter) String() string {
	return fmt.Sprintf("%v", s.Str)
}

func (s *Strconverter) TransformRune(pos int) {
	if isDigit(myrune[pos]) {
		counter++
	} else if isLetter(myrune[pos]) {
		myrune[pos] = singleRune([]rune(strings.ToLower(string(myrune[pos]))))
		counter++
		if counter%3 == 0 {
			myrune[pos] = singleRune([]rune(strings.ToUpper(string(myrune[pos]))))

		}

	}
	s.Str = string(myrune)

}

func (s Strconverter) GetValueAsRuneSlice() []rune {
	r := []rune(s.Str)
	return r
}

func MapString(i Interface) {

	myrune = i.GetValueAsRuneSlice()
	for pos := range myrune {
		i.TransformRune(pos)
	}

}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
