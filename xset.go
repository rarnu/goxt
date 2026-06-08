package goxt

type XSet[T Equalable[T]] []T

func NewXSet[T Equalable[T]]() XSet[T] {
	return XSet[T]{}
}

func NewXSetWithSize[T Equalable[T]](size XInt) XSet[T] {
	return make(XSet[T], 0, size)
}

func NewXSetWithElements[T Equalable[T]](elements ...T) XSet[T] {
	ret := make(XSet[T], 0, len(elements))
	for _, item := range elements {
		if !ret.Contains(item) {
			ret.Add(item)
		}
	}
	return ret
}

func NewXSetWithInit[T Equalable[T]](size XInt, init func(XInt) T) XSet[T] {
	ret := make(XSet[T], 0, size)
	for i := 0; i < int(size); i++ {
		item := init(XInt(i))
		if !ret.Contains(item) {
			ret.Add(item)
		}
	}
	return ret
}

func EmptyXSet[T Equalable[T]]() XSet[T] {
	return make(XSet[T], 0)
}

func (s XSet[T]) Equal(other XSet[T]) XBool {
	if len(s) != len(other) {
		return false
	}
	for _, key := range other {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

func (s XSet[T]) Size() XInt {
	return XInt(len(s))
}

func (s XSet[T]) IsEmpty() XBool {
	return s.Size() == 0
}

func (s XSet[T]) IsNotEmpty() XBool {
	return s.Size() != 0
}

func (s XSet[T]) Contains(element T) XBool {
	if s == nil || len(s) == 0 {
		return false
	}
	for _, item := range s {
		if item.Equal(element) {
			return true
		}
	}
	return false
}

func (s XSet[T]) ContainsAll(elements XList[T]) XBool {
	for _, item := range elements {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s *XSet[T]) Add(element T) XBool {
	if !s.Contains(element) {
		*s = append(*s, element)
		return true
	}
	return false
}

func (s *XSet[T]) Remove(element T) XBool {
	for i, item := range *s {
		if item.Equal(element) {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return true
		}
	}
	return false
}

func (s *XSet[T]) AddAll(elements XList[T]) XBool {
	hasError := false
	for _, item := range elements {
		if !s.Add(item) {
			hasError = true
		}
	}
	return !XBool(hasError)
}

func (s *XSet[T]) RemoveAll(elements XList[T]) XBool {
	hasError := false
	for _, item := range elements {
		if !s.Remove(item) {
			hasError = true
		}
	}
	return !XBool(hasError)
}

func (s *XSet[T]) RetainAll(elements XList[T]) XBool {
	if s.IsEmpty() {
		return false
	}
	ret := make(XSet[T], 0, len(elements))
	for _, v := range *s {
		if elements.Contains(v) {
			if !ret.Contains(v) {
				ret.Add(v)
			}
		}
	}
	if ret.Equal(*s) {
		return false
	}
	*s = ret
	return true
}

func (s *XSet[T]) Clear() {
	*s = EmptyXSet[T]()
}

func (s XSet[T]) IfEmpty(defaultValue func() XSet[T]) XSet[T] {
	if s.IsEmpty() {
		return defaultValue()
	}
	return s
}

func (s *XSet[T]) RemoveAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make(XSet[T], 0, len(*s))
	for _, elem := range *s {
		if !predicate(elem) {
			if !newItems.Contains(elem) {
				newItems.Add(elem)
				result = true
			}
		}
	}
	*s = newItems
	return XBool(result)
}
func (s *XSet[T]) RetainAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make(XSet[T], 0, len(*s))
	for _, elem := range *s {
		if predicate(elem) {
			if !newItems.Contains(elem) {
				newItems.Add(elem)
				result = true
			}
		}
	}
	*s = newItems
	return XBool(result)
}
