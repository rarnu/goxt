package goxt

type XMap[K comparable, V comparable] map[K]V

func NewXMap[K comparable, V comparable]() XMap[K, V] {
	return XMap[K, V]{}
}

func NewXMapWithSize[K comparable, V comparable](size XInt) XMap[K, V] {
	return make(XMap[K, V], size)
}

func NewXMapWithElements[K comparable, V comparable](elements ...XPair[K, V]) XMap[K, V] {
	ret := make(XMap[K, V], len(elements))
	for i := 0; i < len(elements); i++ {
		ret[elements[i].First] = elements[i].Second
	}
	return ret
}

func NewXMapWithInit[K comparable, V comparable](size XInt, init func(XInt) XPair[K, V]) XMap[K, V] {
	ret := make(XMap[K, V], size)
	for i := 0; i < int(size); i++ {
		entry := init(XInt(i))
		ret[entry.First] = entry.Second
	}
	return ret
}

func EmptyXMap[K comparable, V comparable]() XMap[K, V] {
	return make(XMap[K, V])
}

func (m *XMap[K, V]) Size() XInt {
	return XInt(len(*m))
}
func (m *XMap[K, V]) IsEmpty() XBool {
	return len(*m) == 0
}
func (m *XMap[K, V]) IsNotEmpty() XBool {
	return len(*m) != 0
}
func (m *XMap[K, V]) ContainsKey(key K) XBool {
	_, ok := (*m)[key]
	return XBool(ok)
}
func (m *XMap[K, V]) ContainsValue(value V) XBool {
	for _, v := range *m {
		if v == value {
			return true
		}
	}
	return false
}
func (m *XMap[K, V]) Keys() XSet[K] {
	keys := NewXSetWithSize[K](m.Size())
	for k := range *m {
		keys.Add(k)
	}
	return keys
}
func (m *XMap[K, V]) Values() XList[V] {
	values := NewXListWithSize[V](m.Size())
	for _, v := range *m {
		values.Add(v)
	}
	return values
}
func (m *XMap[K, V]) Entries() XSet[XPair[K, V]] {
	ent := NewXSetWithSize[XPair[K, V]](m.Size())
	for k, v := range *m {
		ent.Add(XPair[K, V]{First: k, Second: v})
	}
	return ent
}
func (m *XMap[K, V]) Remove(key K) *V {
	if v, ok := (*m)[key]; ok {
		delete(*m, key)
		return &v
	}
	return nil
}
func (m *XMap[K, V]) PutAll(from XMap[K, V]) {
	for k, v := range from {
		(*m)[k] = v
	}
}
func (m *XMap[K, V]) Clear() {
	*m = make(XMap[K, V])
}

func (m *XMap[K, V]) IfEmpty(defaultValue func() XMap[K, V]) XMap[K, V] {
	if m.IsEmpty() {
		return defaultValue()
	}
	return *m
}

func (m *XMap[K, V]) PutAllPairs(pairs ...XPair[K, V]) {
	for _, p := range pairs {
		(*m)[p.First] = p.Second
	}
}
