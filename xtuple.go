package goxt

import "fmt"

type XPair[A comparable, B comparable] struct {
	First  A
	Second B
}

type XTriple[A comparable, B comparable, C comparable] struct {
	First  A
	Second B
	Third  C
}

func NewPair[A comparable, B comparable](first A, second B) XPair[A, B] {
	return XPair[A, B]{
		First:  first,
		Second: second,
	}
}

func NewTriple[A comparable, B comparable, C comparable](first A, second B, third C) XTriple[A, B, C] {
	return XTriple[A, B, C]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

func (p XPair[A, B]) Equal(other XPair[A, B]) XBool {
	return p.First == other.First && p.Second == other.Second
}

func (t XTriple[A, B, C]) Equal(other XTriple[A, B, C]) XBool {
	return t.First == other.First && t.Second == other.Second && t.Third == other.Third
}

func (p XPair[A, B]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

func (t XTriple[A, B, C]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.First, t.Second, t.Third)
}

func PairToList[A Comparable[A]](p XPair[A, A]) XList[A] {
	return []A{p.First, p.Second}
}

func TripleToList[A Comparable[A]](t XTriple[A, A, A]) XList[A] {
	return []A{t.First, t.Second, t.Third}
}
