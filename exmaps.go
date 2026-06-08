package goxt

func (m XMap[K, V]) MapKeysTo[R Comparable[R]](destination XMap[R, V], transform func(XMapEntry[K, V]) R) XMap[R, V] {
	for _, entry := range m.Entries() {
		r := transform(entry)
		destination[r] = entry.Value
	}
	return destination
}

func (m XMap[K, V]) MapValuesTo[R Comparable[R]](destination XMap[K, R], transform func(XMapEntry[K, V]) R) XMap[K, R] {
	for _, entry := range m.Entries() {
		r := transform(entry)
		destination[entry.Key] = r
	}
	return destination
}

func (m XMap[K, V]) MapKeys[R Comparable[R]](transform func(XMapEntry[K, V]) R) XMap[R, V] {
	dest := make(XMap[R, V], len(m))
	return m.MapKeysTo(dest, transform)
}

func (m XMap[K, V]) MapValues[R Comparable[R]](transform func(XMapEntry[K, V]) R) XMap[K, R] {
	dest := make(XMap[K, R], len(m))
	return m.MapValuesTo(dest, transform)
}

func (m XMap[K, V]) FilterKeys(predicate func(K) XBool) XMap[K, V] {
	dest := make(XMap[K, V], len(m))
	for k, v := range m {
		if predicate(k) {
			dest[k] = v
		}
	}
	return dest
}

func (m XMap[K, V]) FilterValues(predicate func(V) XBool) XMap[K, V] {
	dest := make(XMap[K, V], len(m))
	for k, v := range m {
		if predicate(v) {
			dest[k] = v
		}
	}
	return dest
}

func (m XMap[K, V]) FilterTo(destination XMap[K, V], predicate func(XMapEntry[K, V]) XBool) XMap[K, V] {
	for _, entry := range m.Entries() {
		if predicate(entry) {
			destination[entry.Key] = entry.Value
		}
	}
	return destination
}

func (m XMap[K, V]) Filter(predicate func(XMapEntry[K, V]) XBool) XMap[K, V] {
	dest := make(XMap[K, V], len(m))
	return m.FilterTo(dest, predicate)
}

func (m XMap[K, V]) FilterNotTo(destination XMap[K, V], predicate func(XMapEntry[K, V]) XBool) XMap[K, V] {
	for _, entry := range m.Entries() {
		if !predicate(entry) {
			destination[entry.Key] = entry.Value
		}
	}
	return destination
}

func (m XMap[K, V]) FilterNot(predicate func(XMapEntry[K, V]) XBool) XMap[K, V] {
	dest := make(XMap[K, V], len(m))
	return m.FilterNotTo(dest, predicate)
}

func (m XMap[K, V]) ToList() XList[XMapEntry[K, V]] {
	ret := make(XList[XMapEntry[K, V]], m.Size())
	for _, entry := range m.Entries() {
		ret.Add(entry)
	}
	return ret
}