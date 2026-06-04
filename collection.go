package goxt

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
}

type List[T comparable] interface {
	Collection[T]
	IndexOf(element T) XInt
	LastIndexOf(element T) XInt
	SubList(fromIndex XInt, toIndex XInt) List[T]
	AddAt(index XInt, element T) XBool
	RemoveAt(index XInt) *T
}

type Set[T comparable] interface {
	Collection[T]
}

type MapEntry[K comparable, V comparable] struct{
	Key K
	Value V
}

type Map[K comparable, V comparable] interface {
	Size() XInt
	IsEmpty() XBool
	IsNotEmpty() XBool
	ContainsKey(key K) XBool
	ContainsValue(value V) XBool
	Keys() Set[K]
	Values() Collection[V]
	Entries() Set[MapEntry[K, V]]
	Remove(key K) *V
	PutAll(from Map[K, V])
	Clear()
}