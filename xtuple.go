package goxt

import "fmt"

type XPair[A Equalable[A], B Equalable[B]] struct {
	First  A
	Second B
}

type XTriple[A Equalable[A], B Equalable[B], C Equalable[C]] struct {
	First  A
	Second B
	Third  C
}

func NewPair[A Equalable[A], B Equalable[B]](first A, second B) XPair[A, B] {
	return XPair[A, B]{
		First:  first,
		Second: second,
	}
}

func NewTriple[A Equalable[A], B Equalable[B], C Equalable[C]](first A, second B, third C) XTriple[A, B, C] {
	return XTriple[A, B, C]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

func (p XPair[A, B]) Equal(other XPair[A, B]) XBool {
	return p.First.Equal(other.First) && p.Second.Equal(other.Second)
}

func (t XTriple[A, B, C]) Equal(other XTriple[A, B, C]) XBool {
	return t.First.Equal(other.First) && t.Second.Equal(other.Second) && t.Third.Equal(other.Third)
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
