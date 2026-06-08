package goxt

func XPairedListToMap[K Comparable[K], V Equalable[V]](l XList[XMapEntry[K, V]]) XMap[K, V] {
	if l.IsEmpty() {
		return XMap[K, V]{}
	}
	ret := make(XMap[K, V], l.Size())
	for _, item := range l {
		ret[item.Key] = item.Value
	}
	return ret
}
