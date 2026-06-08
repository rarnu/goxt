package goxt

import "math"

func XEntryListToMap[K Comparable[K], V Equalable[V]](l XList[XMapEntry[K, V]]) XMap[K, V] {
	if l.IsEmpty() {
		return XMap[K, V]{}
	}
	ret := make(XMap[K, V], l.Size())
	for _, item := range l {
		ret[item.Key] = item.Value
	}
	return ret
}

func XListMax[T NumbersConstraints[T]](l XList[T]) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	for _, item := range l[1:] {
		if item > m {
			m = item
		}
	}
	return m
}

func XListMaxOrNil[T NumbersConstraints[T]](l XList[T]) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	for _, item := range l[1:] {
		if item > m {
			m = item
		}
	}
	return &m
}

func XListMaxBy[T NumbersConstraints[T], R NumbersConstraints[T]](l XList[T], selector func(T) R) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	mv := selector(m)
	for _, e := range l[1:] {
		v := selector(e)
		if mv < v {
			m = e
			mv = v
		}
	}
	return m
}

func XListMaxByOrNil[T NumbersConstraints[T], R NumbersConstraints[T]](l XList[T], selector func(T) R) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	mv := selector(m)
	for _, e := range l[1:] {
		v := selector(e)
		if mv < v {
			m = e
			mv = v
		}
	}
	return &m
}

func XListMaxOf[T NumbersConstraints[T]](l XList[T], selector func(T) T) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	mv := selector(l[0])
	for _, item := range l[1:] {
		v := selector(item)
		if v > mv {
			mv = v
		}
	}
	return mv
}

func XListMaxOfOrNil[T NumbersConstraints[T]](l XList[T], selector func(T) T) *T {
	if l.IsEmpty() {
		return nil
	}
	mv := selector(l[0])
	for _, item := range l[1:] {
		v := selector(item)
		if v > mv {
			mv = v
		}
	}
	return &mv
}

func XListMin[T NumbersConstraints[T]](l XList[T]) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	for _, item := range l[1:] {
		if item < m {
			m = item
		}
	}
	return m
}

func XListMinOrNil[T NumbersConstraints[T]](l XList[T]) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	for _, item := range l[1:] {
		if item < m {
			m = item
		}
	}
	return &m
}

func XListMinBy[T NumbersConstraints[T], R NumbersConstraints[T]](l XList[T], selector func(T) R) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	mv := selector(m)
	for _, e := range l[1:] {
		v := selector(e)
		if mv > v {
			m = e
			mv = v
		}
	}
	return m
}

func XListMinByOrNil[T NumbersConstraints[T], R NumbersConstraints[T]](l XList[T], selector func(T) R) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	mv := selector(m)
	for _, e := range l[1:] {
		v := selector(e)
		if mv > v {
			m = e
			mv = v
		}
	}
	return &m
}

func XListMinOf[T NumbersConstraints[T]](l XList[T], selector func(T) T) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	mv := selector(l[0])
	for _, item := range l[1:] {
		v := selector(item)
		if v < mv {
			mv = v
		}
	}
	return mv
}

func XListMinOfOrNil[T NumbersConstraints[T]](l XList[T], selector func(T) T) *T {
	if l.IsEmpty() {
		return nil
	}
	mv := selector(l[0])
	for _, item := range l[1:] {
		v := selector(item)
		if v < mv {
			mv = v
		}
	}
	return &mv
}

func XListSumOf[T NumbersConstraints[T]](l XList[T], selector func(T) T) T {
	sm := T(0)
	for _, item := range l {
		sm += selector(item)
	}
	return sm
}

func XListAverage[T NumbersConstraints[T]](l XList[T]) XFloat64 {
	sm := 0.0
	for _, item := range l {
		sm = sm + float64(item)
	}
	if l.Size() == 0 {
		return XFloat64(math.NaN())
	}
	return XFloat64(sm / float64(l.Size()))
}

func XListSum[T NumbersConstraints[T]](l XList[T]) T {
	sm := T(0)
	for _, item := range l {
		sm = sm + item
	}
	return sm
}