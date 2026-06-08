package goxt

type XMapEntry[K Comparable[K], V Equalable[V]] struct {
	Key K
	Value V
}

type XMap[K Comparable[K], V Equalable[V]] map[K]V

func NewXMap[K Comparable[K], V Equalable[V]]() XMap[K, V] {
	return XMap[K, V]{}
}

func NewXMapWithSize[K Comparable[K], V Equalable[V]](size XInt) XMap[K, V] {
	return make(XMap[K, V], size)
}

func NewXMapWithElements[K Comparable[K], V Equalable[V]](elements ...XMapEntry[K, V]) XMap[K, V] {
	ret := make(XMap[K, V], len(elements))
	for i := 0; i < len(elements); i++ {
		ret[elements[i].Key] = elements[i].Value
	}
	return ret
}

func NewXMapWithInit[K Comparable[K], V Equalable[V]](size XInt, init func(XInt) XMapEntry[K, V]) XMap[K, V] {
	ret := make(XMap[K, V], size)
	for i := 0; i < int(size); i++ {
		entry := init(XInt(i))
		ret[entry.Key] = entry.Value
	}
	return ret
}

func EmptyXMap[K Comparable[K], V Equalable[V]]() XMap[K, V] {
	return make(XMap[K, V])
}

func (m XMap[K, V]) Equal(other XMap[K, V]) XBool {
	if m.Size() != other.Size() {
		return false
	}
	for k, v := range m {
		if v0, ok := other[k]; !ok || bool(!v0.Equal(v)) {
			return false
		}
	}
	return true
}

func (e XMapEntry[K, V]) Equal(other XMapEntry[K, V]) XBool {
	return e.Key == other.Key && e.Value.Equal(other.Value)
}

func (m XMap[K, V]) Size() XInt {
	return XInt(len(m))
}
func (m XMap[K, V]) IsEmpty() XBool {
	return len(m) == 0
}
func (m XMap[K, V]) IsNotEmpty() XBool {
	return len(m) != 0
}
func (m XMap[K, V]) ContainsKey(key K) XBool {
	_, ok := m[key]
	return XBool(ok)
}
func (m XMap[K, V]) ContainsValue(value V) XBool {
	for _, v := range m {
		if v.Equal(value) {
			return true
		}
	}
	return false
}
func (m XMap[K, V]) Keys() XSet[K] {
	keys := NewXSetWithSize[K](m.Size())
	for k := range m {
		keys.Add(k)
	}
	return keys
}
func (m XMap[K, V]) Values() XList[V] {
	values := NewXListWithSize[V](m.Size())
	for _, v := range m {
		values.Add(v)
	}
	return values
}
func (m XMap[K, V]) Entries() XSet[XMapEntry[K, V]] {
	ent := NewXSetWithSize[XMapEntry[K, V]](m.Size())
	for k, v := range m {
		ent.Add(XMapEntry[K, V]{Key: k, Value: v})
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

func (m *XMap[K, V]) RemoveAll(keys XSet[K]) {
	for _, k := range keys {
		delete(*m, k)
	}
}

func (m *XMap[K, V]) Clear() {
	*m = make(XMap[K, V])
}

func (m XMap[K, V]) IfEmpty(defaultValue func() XMap[K, V]) XMap[K, V] {
	if m.IsEmpty() {
		return defaultValue()
	}
	return m
}

func (m *XMap[K, V]) PutAllEntries(entries ...XMapEntry[K, V]) {
	for _, p := range entries {
		(*m)[p.Key] = p.Value
	}
}
