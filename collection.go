package goxt

type Collection[T Equalable[T]] interface {
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

type List[T Equalable[T]] interface {
	Collection[T]
	IndexOf(element T) XInt
	LastIndexOf(element T) XInt
	Insert(index XInt, element T) XBool
	RemoveAt(index XInt) *T
	SubList(fromIndex XInt, toIndex XInt) XList[T]
}

type Set[T Equalable[T]] interface {
	Collection[T]
}

type Map[K Comparable[K], V Equalable[V]] interface {
	Size() XInt
	IsEmpty() XBool
	IsNotEmpty() XBool
	ContainsKey(key K) XBool
	ContainsValue(value V) XBool
	Keys() XSet[K]
	Values() XList[V]
	Entries() XSet[XMapEntry[K, V]]
	Remove(key K) *V
	PutAll(from XMap[K, V])
	RemoveAll(keys XSet[K])
	Clear()
}

var (
	// XLIST
	_ Collection[Nothing] = (*XList[Nothing])(nil)
	_ List[Nothing]       = (*XList[Nothing])(nil)

	// XSET
	_ Collection[Nothing] = (*XSet[Nothing])(nil)
	_ Set[Nothing]        = (*XSet[Nothing])(nil)

	_ Map[Nothing, Nothing] = (*XMap[Nothing, Nothing])(nil)
)
