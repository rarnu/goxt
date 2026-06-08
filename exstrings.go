package goxt

import "math/rand"

func (s XString) First() XRune {
	if s.IsEmpty() {
		panic("String is empty.")
	}
	return s.RuneAt(0)
}

func (s XString) FirstOrNil() *XRune {
	if s.IsEmpty() {
		return nil
	}
	return new(s.RuneAt(0))
}

func (s XString) FirstWithPredicate(predicate func(XRune) XBool) XRune {
	for _, item := range []XRune(s) {
		if predicate(item) {
			return item
		}
	}
	panic("String contains no element matching the predicate.")
}

func (s XString) FirstWithPredicateOrNil(predicate func(XRune) XBool) *XRune {
	for _, item := range []XRune(s) {
		if predicate(item) {
			return new(item)
		}
	}
	return nil
}

func (s XString) Last() XRune {
	if s.IsEmpty() {
		panic("String is empty.")
	}
	return s.RuneAt(s.LastIndex())
}

func (s XString) LastOrNil() *XRune {
	if s.IsEmpty() {
		return nil
	}
	return new(s.RuneAt(s.LastIndex()))
}

func (s XString) LastWithPredicate(predicate func(XRune) XBool) XRune {
	rs := []XRune(s)
	for index := len(rs) - 1; index >= 0; index-- {
		item := rs[index]
		if predicate(item) {
			return item
		}
	}
	panic("String contains no element matching the predicate.")
}

func (s XString) LastWithPredicateOrNil(predicate func(XRune) XBool) *XRune {
	rs := []XRune(s)
	for index := len(rs) - 1; index >= 0; index-- {
		item := rs[index]
		if predicate(item) {
			return new(item)
		}
	}
	return nil
}

func (s XString) IndexOfFirst(predicate func(XRune) XBool) XInt {
	for index, item := range []XRune(s) {
		if predicate(item) {
			return XInt(index)
		}
	}
	return -1
}

func (s XString) IndexOfLast(predicate func(XRune) XBool) XInt {
	rs := []XRune(s)
	for index := len(rs) - 1; index >= 0; index-- {
		item := rs[index]
		if predicate(item) {
			return XInt(index)
		}
	}
	return -1
}

func (s XString) Random() XRune {
	if s.IsEmpty() {
		panic("String is empty.")
	}
	return s.RuneAt(XInt(rand.Intn(int(s.RuneCount()))))
}

func (s XString) RandomOrNil() *XRune {
	if s.IsEmpty() {
		return nil
	}
	return new(s.RuneAt(XInt(rand.Intn(int(s.RuneCount())))))
}

func (s XString) Drop(n XInt) XString {
	if n < 0 {
		panic("Negative argument.")
	}
	if n >= s.RuneCount() {
		return ""
	}
	return XString([]rune(s)[n:])
}

func (s XString) DropLast(n XInt) XString {
	if n < 0 {
		panic("Negative argument.")
	}
	if n >= s.RuneCount() {
		return ""
	}
	return XString([]rune(s)[:s.RuneCount()-n])
}

func (s XString) Take(n XInt) XString {
	if n < 0 {
		panic("Negative argument.")
	}
	if n >= s.RuneCount() {
		return s
	}
	return XString([]rune(s)[:n])
}

func (s XString) TakeLast(n XInt) XString {
	if n < 0 {
		panic("Negative argument.")
	}
	if n >= s.RuneCount() {
		return s
	}
	return XString([]rune(s)[s.RuneCount()-n:])
}

func (s XString) FilterTo(destination XString, predicate func(XRune) XBool) XString {
	for _, item := range []XRune(s) {
		if predicate(item) {
			destination += XString(item)
		}
	}
	return destination
}

func (s XString) Filter(predicate func(XRune) XBool) XString {
	return s.FilterTo("", predicate)
}

func (s XString) FilterIndexedTo(destination XString, predicate func(XInt, XRune) XBool) XString {
	for index, item := range []XRune(s) {
		if predicate(XInt(index), item) {
			destination += XString(item)
		}
	}
	return destination
}

func (s XString) FilterIndexed(predicate func(XInt, XRune) XBool) XString {
	return s.FilterIndexedTo("", predicate)
}

func (s XString) FilterNotTo(destination XString, predicate func(XRune) XBool) XString {
	for _, item := range []XRune(s) {
		if !predicate(item) {
			destination += XString(item)
		}
	}
	return destination
}

func (s XString) FilterNot(predicate func(XRune) XBool) XString {
	return s.FilterNotTo("", predicate)
}

func (s XString) FilterNotIndexedTo(destination XString, predicate func(XInt, XRune) XBool) XString {
	for index, item := range []XRune(s) {
		if !predicate(XInt(index), item) {
			destination += XString(item)
		}
	}
	return destination
}

func (s XString) FilterNotIndexed(predicate func(XInt, XRune) XBool) XString {
	return s.FilterNotIndexedTo("", predicate)
}

func (s XString) Reversed() XString {
	rs := []XRune(s)
	n := len(rs)
	reversed := make([]XRune, 0, n)
	for i, elem := range rs {
		reversed[n-1-i] = elem
	}
	return XString(reversed)
}

