package main

// https://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite

import "fmt"

type S struct {
	a, b int
}

// String implements the fmt.Stringer interface
func (s *S) String() string {
	return fmt.Sprintf("%s", s) // Sprintf will call s.String()
}

func main() {
	s := &S{a: 1, b: 2}
	fmt.Println(s)
}
