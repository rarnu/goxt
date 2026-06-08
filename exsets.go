package goxt

import (
	"math/rand"
	"time"
)

func (s XSet[T]) Random() T {
	if s.IsEmpty() {
		panic("Collection is empty.")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return s[rand.Intn(len(s))]
}

func (s XSet[T]) RandomOrNil() *T {
	if s.IsEmpty() {
		return nil
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &s[rand.Intn(len(s))]
}

func (s XSet[T]) FilterTo(destination XSet[T], predicate func(T) XBool) XSet[T] {
	for _, item := range s {
		if predicate(item) {
			destination.Add(item)
		}
	}
	return destination
}

func (s XSet[T]) Filter(predicate func(T) XBool) XSet[T] {
	return s.FilterTo(NewXSet[T](), predicate)
}

func (s XSet[T]) FilterNotTo(destination XSet[T], predicate func(T) XBool) XSet[T] {
	for _, item := range s {
		if !predicate(item) {
			destination.Add(item)
		}
	}
	return destination
}

func (s XSet[T]) FilterNot(predicate func(T) XBool) XSet[T] {
	return s.FilterNotTo(NewXSet[T](), predicate)
}

func (s XSet[T]) FilterIndexedTo(destination XSet[T], predicate func(index XInt, item T) XBool) XSet[T] {
	for index, item := range s {
		if predicate(XInt(index), item) {
			destination.Add(item)
		}
	}
	return destination
}

func (s XSet[T]) FilterIndexed(predicate func(index XInt, item T) XBool) XSet[T] {
	return s.FilterIndexedTo(NewXSet[T](), predicate)
}

func (s XSet[T]) AssociateTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], transform func(T) XMapEntry[K, V]) XMap[K, V] {
	for _, item := range s {
		pair := transform(item)
		destination[pair.Key] = pair.Value
	}
	return destination
}

func (s XSet[T]) Associate[K Comparable[K], V Equalable[V]](transform func(T) XMapEntry[K, V]) XMap[K, V] {
	dest := make(XMap[K, V], s.Size())
	return s.AssociateTo(dest, transform)
}

func (s XSet[T]) AssociateByWithValueTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], selector func(T) K, transform func(T) V) XMap[K, V] {
	for _, item := range s {
		key := selector(item)
		value := transform(item)
		destination[key] = value
	}
	return destination
}

func (s XSet[T]) AssociateByWithValue[K Comparable[K], V Equalable[V]](selector func(T) K, transform func(T) V) XMap[K, V] {
	dest := make(XMap[K, V], s.Size())
	return s.AssociateByWithValueTo(dest, selector, transform)
}

func (s XSet[T]) Map[R Equalable[R]](transform func(T) R) XList[R] {
	return s.MapTo(NewXList[R](), transform)
}

func (l XSet[T]) MapTo[R Equalable[R]](destination XList[R], transform func(T) R) XList[R] {
	for _, item := range l {
		destination.Add(transform(item))
	}
	return destination
}

func (s XSet[T]) Intersect(other XSet[T]) XSet[T] {
	st := NewXSet[T]()
	for _, item := range s {
		if other.Contains(item) {
			st.Add(item)
		}
	}
	return st
}

func (s XSet[T]) Subtract(other XSet[T]) XSet[T] {
	st := NewXSet[T]()
	for _, item := range s {
		if !other.Contains(item) {
			st.Add(item)
		}
	}
	return st
}

func (s XSet[T]) Union(other XSet[T]) XSet[T] {
	st := s
	st.AddAll(other.ToList())
	return st
}

func (s XSet[T]) All(predicate func(T) XBool) XBool {
	if s.IsEmpty() {
		return true
	}
	for _, item := range s {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s XSet[T]) Any(predicate func(T) XBool) XBool {
	if s.IsEmpty() {
		return false
	}
	for _, item := range s {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s XSet[T]) Count(predicate func(T) XBool) XInt {
	count := 0
	for _, item := range s {
		if predicate(item) {
			count++
		}
	}
	return XInt(count)
}

func (s XSet[T]) None(predicate func(T) XBool) XBool {
	if s.IsEmpty() {
		return true
	}
	for _, item := range s {
		if predicate(item) {
			return false
		}
	}
	return true
}

func (s XSet[T]) ToList() XList[T] {
	ret := make(XList[T], s.Size())
	for _, item := range s {
		ret.Add(item)
	}
	return ret
}

func (s XSet[T]) MaxWith(comparator func(T, T) XInt) T {
	if s.IsEmpty() {
		panic("Collection is empty.")
	}
	m := s[0]
	for _, item := range s[1:] {
		if comparator(item, m) > 0 {
			m = item
		}
	}
	return m
}

func (s XSet[T]) MaxWithOrNil(comparator func(T, T) XInt) *T {
	if s.IsEmpty() {
		return nil
	}
	m := s[0]
	for _, item := range s[1:] {
		if comparator(item, m) > 0 {
			m = item
		}
	}
	return &m
}

func (s XSet[T]) MinWith(comparator func(T, T) XInt) T {
	if s.IsEmpty() {
		panic("Collection is empty.")
	}
	m := s[0]
	for _, item := range s[1:] {
		if comparator(item, m) < 0 {
			m = item
		}
	}
	return m
}

func (s XSet[T]) MinWithOrNil(comparator func(T, T) XInt) *T {
	if s.IsEmpty() {
		return nil
	}
	m := s[0]
	for _, item := range s[1:] {
		if comparator(item, m) < 0 {
			m = item
		}
	}
	return &m
}