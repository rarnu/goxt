package goxt

import (
	"math"
	"math/rand"
)

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

func (s XString) AssociateTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], transform func(xRune XRune) XMapEntry[K, V]) XMap[K, V] {
	for _, item := range []XRune(s) {
		ele := transform(item)
		destination[ele.Key] = ele.Value
	}
	return destination
}

func (s XString) Associate[K Comparable[K], V Equalable[V]](transform func(xRune XRune) XMapEntry[K, V]) XMap[K, V] {
	dest := make(XMap[K, V], s.RuneCount())
	return s.AssociateTo(dest, transform)
}

func (s XString) AssociateByTo[K Comparable[K]](destination XMap[K, XRune], selector func(XRune) K) XMap[K, XRune] {
	for _, item := range []XRune(s) {
		key := selector(item)
		destination[key] = item
	}
	return destination
}

func (s XString) AssociateBy[K Comparable[K]](selector func(XRune) K) XMap[K, XRune] {
	dest := make(XMap[K, XRune], s.RuneCount())
	return s.AssociateByTo(dest, selector)
}

func (s XString) AssociateByWithValueTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], selector func(XRune) K, valueSelector func(XRune) V) XMap[K, V] {
	for _, item := range []XRune(s) {
		key := selector(item)
		value := valueSelector(item)
		destination[key] = value
	}
	return destination
}

func (s XString) AssociateByWithValue[K Comparable[K], V Equalable[V]](selector func(XRune) K, valueSelector func(XRune) V) XMap[K, V] {
	dest := make(XMap[K, V], s.RuneCount())
	return s.AssociateByWithValueTo(dest, selector, valueSelector)
}

func (s XString) AssociateWithTo[V Equalable[V]](destinatioon XMap[XRune, V], valueSelector func(XRune) V) XMap[XRune, V] {
	for _, item := range []XRune(s) {
		value := valueSelector(item)
		destinatioon[item] = value
	}
	return destinatioon
}

func (s XString) AssociateWith[V Equalable[V]](valueSelector func(XRune) V) XMap[XRune, V] {
	dest := make(XMap[XRune, V], s.RuneCount())
	return s.AssociateWithTo(dest, valueSelector)
}

func (s XString) GroupByTo[K Comparable[K]](destination XMap[K, XList[XRune]], selector func(XRune) K) XMap[K, XList[XRune]] {
	for _, item := range []XRune(s) {
		key := selector(item)
		list := destination.GetOrPut(key, func() XList[XRune] {
			return NewXList[XRune]()
		})
		list.Add(item)
	}
	return destination
}

func (s XString) GroupBy[K Comparable[K]](selector func(XRune) K) XMap[K, XList[XRune]] {
	return s.GroupByTo(make(XMap[K, XList[XRune]]), selector)
}

func (s XString) GroupByWithValueTo[K Comparable[K], V Equalable[V]](destination XMap[K, XList[V]], selector func(XRune) K, valueTransform func(XRune) V) XMap[K, XList[V]] {
	for _, item := range []XRune(s) {
		key := selector(item)
		list := destination.GetOrPut(key, func() XList[V] {
			return NewXList[V]()
		})
		list.Add(valueTransform(item))
	}
	return destination
}

func (s XString) GroupByWithValue[K Comparable[K], V Equalable[V]](selector func(XRune) K, valueTransform func(XRune) V) XMap[K, XList[V]] {
	return s.GroupByWithValueTo(make(XMap[K, XList[V]]), selector, valueTransform)
}

func (s XString) MapTo[R Equalable[R]](destination XList[R], transform func(XRune) R) XList[R] {
	for _, item := range []XRune(s) {
		destination.Add(transform(item))
	}
	return destination
}

func (s XString) Map[R Equalable[R]](transform func(XRune) R) XList[R] {
	dest := make(XList[R], 0, s.RuneCount())
	return s.MapTo(dest, transform)
}

func (s XString) MapIndexedTo[R Equalable[R]](destination XList[R], transform func(XInt, XRune) R) XList[R] {
	for index, item := range []XRune(s) {
		destination.Add(transform(XInt(index), item))
	}
	return destination
}

func (s XString) MapIndexed[R Equalable[R]](transform func(XInt, XRune) R) XList[R] {
	dest := make(XList[R], 0, s.RuneCount())
	return s.MapIndexedTo(dest, transform)
}

func (s XString) All(predicate func(XRune) XBool) XBool {
	if s.IsEmpty() {
		return false
	}
	for _, item := range []XRune(s) {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s XString) Any(predicate func(XRune) XBool) XBool {
	if s.IsEmpty() {
		return false
	}
	for _, item := range []XRune(s) {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s XString) None(predicate func(XRune) XBool) XBool {
	if s.IsEmpty() {
		return true
	}
	for _, item := range []XRune(s) {
		if predicate(item) {
			return false
		}
	}
	return true
}

func (s XString) Count(predicate func(XRune) XBool) XInt {
	count := XInt(0)
	for _, item := range []XRune(s) {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (s XString) Fold[R Equalable[R]](initial R, operation func(R, XRune) R) R {
	acc := initial
	for _, item := range []XRune(s) {
		acc = operation(acc, item)
	}
	return acc
}
func (s XString) FoldIndexed[R Equalable[R]](initial R, operation func(XInt, R, XRune) R) R {
	acc := initial
	for index, item := range []XRune(s) {
		acc = operation(XInt(index), acc, item)
	}
	return acc
}

func (s XString) Reduce(operation func(XRune, XRune) XRune) XRune {
	if s.IsEmpty() {
		panic("Empty string can't be reduced.")
	}
	rs := []XRune(s)
	acc := rs[0]
	for _, item := range rs[1:] {
		acc = operation(acc, item)
	}
	return acc
}

func (s XString) ReduceOrNil(operation func(XRune, XRune) XRune) *XRune {
	if s.IsEmpty() {
		return nil
	}
	rs := []XRune(s)
	acc := rs[0]
	for _, item := range rs[1:] {
		acc = operation(acc, item)
	}
	return &acc
}

func (s XString) ReduceIndexed(operation func(XInt, XRune, XRune) XRune) XRune {
	if s.IsEmpty() {
		panic("Empty string can't be reduced.")
	}
	rs := []XRune(s)
	acc := rs[0]
	for index, item := range rs[1:] {
		acc = operation(XInt(index+1), acc, item)
	}
	return acc
}

func (s XString) ReduceIndexedOrNil(operation func(XInt, XRune, XRune) XRune) *XRune {
	if s.IsEmpty() {
		return nil
	}
	rs := []XRune(s)
	acc := rs[0]
	for index, item := range rs[1:] {
		acc = operation(XInt(index+1), acc, item)
	}
	return &acc
}

func (s XString) RunningFold[R Equalable[R]](initial R, operation func(R, XRune) R) XList[R] {
	if s.IsEmpty() {
		return NewXListWithElements(initial)
	}
	rs := []XRune(s)
	dest := make(XList[R], 0, len(rs)+1)
	dest.Add(initial)
	acc := initial
	for _, item := range rs {
		acc = operation(acc, item)
		dest.Add(acc)
	}
	return dest
}

func (s XString) RunningFoldIndexed[R Equalable[R]](initial R, operation func(XInt, R, XRune) R) XList[R] {
	if s.IsEmpty() {
		return NewXListWithElements(initial)
	}
	rs := []XRune(s)
	dest := make(XList[R], 0, len(rs)+1)
	dest.Add(initial)
	acc := initial
	for index, item := range rs {
		acc = operation(XInt(index), acc, item)
		dest.Add(acc)
	}
	return dest
}

func (s XString) RunningReduce(operation func(XRune, XRune) XRune) XList[XRune] {
	if s.IsEmpty() {
		return NewXList[XRune]()
	}
	rs := []XRune(s)
	acc := rs[0]
	dest := make(XList[XRune], 0, len(rs))
	dest.Add(acc)
	for _, item := range rs[1:] {
		acc = operation(acc, item)
		dest.Add(acc)
	}
	return dest
}

func (s XString) RunningReduceIndexed(operation func(XInt, XRune, XRune) XRune) XList[XRune] {
	if s.IsEmpty() {
		return NewXList[XRune]()
	}
	rs := []XRune(s)
	acc := rs[0]
	dest := make(XList[XRune], 0, len(rs))
	dest.Add(acc)
	for index, item := range rs[1:] {
		acc = operation(XInt(index+1), acc, item)
		dest.Add(acc)
	}
	return dest
}

func (s XString) Partition(predicate func(XRune) XBool) XPair[XString, XString] {
	first := XString("")
	second := XString("")
	for _, item := range []XRune(s) {
		if predicate(item) {
			first += XString(item)
		} else {
			second += XString(item)
		}
	}
	return XPair[XString, XString]{First: first, Second: second}
}

func (s XString) ZipWith[V Equalable[V]](other XString, transform func(XRune, XRune) V) XList[V] {
	minLen := int64(math.Min(float64(s.RuneCount()), float64(other.RuneCount())))
	list := make(XList[V], 0, minLen)
	rsThis := []XRune(s)
	rsOther := []XRune(other)
	for index, item := range rsThis {
		list.Add(transform(item, rsOther[index]))
	}
	return list
}

func (s XString) Zip(other XString) XList[XPair[XRune, XRune]] {
	return s.ZipWith(other, func(c1 XRune, c2 XRune) XPair[XRune, XRune] {
		return XPair[XRune, XRune]{First: c1, Second: c2}
	})
}

func (s XString) ZipNextWith[R Equalable[R]](transform func(XRune, XRune) R) XList[R] {
	size := s.RuneCount() - 1
	if size < 1 {
		return NewXList[R]()
	}
	dest := make(XList[R], 0, size)
	rs := []XRune(s)
	for index := range rs {
		dest.Add(transform(rs[index], rs[index+1]))
	}
	return dest
}

func (s XString) ZipNext() XList[XPair[XRune, XRune]] {
	return s.ZipNextWith(func(a XRune, b XRune) XPair[XRune, XRune] {
		return XPair[XRune, XRune]{First: a, Second: b}
	})
}
