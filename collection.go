package goxt

import "go/types"

type Collection[T comparable] interface {
	Size() XInt
	IsEmpty() XBool
	IsNotEmpty() XBool
	Contains(element T) XBool
	ContainsAll(elements XList[T]) XBool
	Add(element T) XBool
	Remove(element T) XBool
	AddAll(elements XList[T]) XBool
	RemoveAll(elements XList[T]) XBool
	RetainAll(elements XList[T]) XBool
	Clear()
	RemoveAllWithPredicate(predicate func(element T) XBool) XBool
	RetainAllWithPredicate(predicate func(element T) XBool) XBool
}

type List[T comparable] interface {
	Collection[T]
	IndexOf(element T) XInt
	LastIndexOf(element T) XInt
	Insert(index XInt, element T) XBool
	RemoveAt(index XInt) *T
	SubList(fromIndex XInt, toIndex XInt) XList[T]
}

type Set[T comparable] interface {
	Collection[T]
}

type Map[K comparable, V comparable] interface {
	Size() XInt
	IsEmpty() XBool
	IsNotEmpty() XBool
	ContainsKey(key K) XBool
	ContainsValue(value V) XBool
	Keys() XSet[K]
	Values() XList[V]
	Entries() XSet[XPair[K, V]]
	Remove(key K) *V
	PutAll(from XMap[K, V])
	Clear()
}

var (
	// XLIST
	_ Collection[types.Nil] = (*XList[types.Nil])(nil)
	_ List[types.Nil]       = (*XList[types.Nil])(nil)

	// XSET
	_ Collection[types.Nil] = (*XSet[types.Nil])(nil)
	_ Set[types.Nil]        = (*XSet[types.Nil])(nil)

	_ Map[types.Nil, types.Nil] = (*XMap[types.Nil, types.Nil])(nil)
)
