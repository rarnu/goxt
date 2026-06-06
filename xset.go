package goxt

type XSet[T comparable] map[T]bool

func NewXSet[T comparable]() XSet[T] {
	return XSet[T]{}
}

func NewXSetWithSize[T comparable](size XInt) XSet[T] {
	return make(XSet[T], size)
}

func NewXSetWithElements[T comparable](elements ...T) XSet[T] {
	ret := make(XSet[T], len(elements))
	for _, item := range elements {
		ret[item] = true
	}
	return ret
}

func NewXSetWithInit[T comparable](size XInt, init func(XInt) T) XSet[T] {
	ret := make(XSet[T], size)
	for i := 0; i < int(size); i++ {
		ret[init(XInt(i))] = true
	}
	return ret
}

func EmptyXSet[T comparable]() XSet[T] {
	return make(XSet[T])
}

func (s *XSet[T]) Size() XInt {
	return XInt(len(*s))
}

func (s *XSet[T]) IsEmpty() XBool {
	return s.Size() == 0
}

func (s *XSet[T]) IsNotEmpty() XBool {
	return s.Size() != 0
}

func (s *XSet[T]) Contains(element T) XBool {
	_, ok := (*s)[element]
	return XBool(ok)
}

func (s *XSet[T]) ContainsAll(elements XList[T]) XBool {
	for _, item := range elements {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s *XSet[T]) Add(element T) XBool {
	(*s)[element] = true
	return true
}

func (s *XSet[T]) Remove(element T) XBool {
	if _, ok := (*s)[element]; ok {
		delete(*s, element)
		return true
	}
	return false
}

func (s *XSet[T]) AddAll(elements XList[T]) XBool {
	for _, item := range elements {
		(*s)[item] = true
	}
	return true
}

func (s *XSet[T]) RemoveAll(elements XList[T]) XBool {
	for _, item := range elements {
		if _, ok := (*s)[item]; ok {
			delete(*s, item)
		}
	}
	return true
}

func (s *XSet[T]) RetainAll(elements XList[T]) XBool {
	if s == nil || len(*s) == 0 {
		return false
	}
	if elements == nil || len(elements) == 0 {
		// 清空当前集合，因为没有任何元素需要保留
		*s = make(XSet[T])
		return true
	}
	// 创建一个临时集合来存储需要保留的元素（去重）
	retainSet := make(XSet[T])
	for _, elem := range elements {
		retainSet[elem] = true
	}
	// 记录是否有变化
	changed := false
	// 遍历当前集合，删除不在 retainSet 中的元素
	for key := range *s {
		if _, exists := retainSet[key]; !exists {
			delete(*s, key)
			changed = true
		}
	}
	if changed {
		return true
	}
	return false
}

func (s *XSet[T]) Clear() {
	*s = EmptyXSet[T]()
}

func (s *XSet[T]) IfEmpty(defaultValue func() XSet[T]) XSet[T] {
	if s.IsEmpty() {
		return defaultValue()
	}
	return *s
}

func (s *XSet[T]) RemoveAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make(XSet[T], len(*s))
	for elem := range *s {
		if !predicate(elem) {
			newItems[elem] = true
			result = true
		}
	}
	*s = newItems
	return XBool(result)
}
func (s *XSet[T]) RetainAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make(XSet[T], len(*s))
	for elem := range *s {
		if predicate(elem) {
			newItems[elem] = true
			result = true
		}
	}
	*s = newItems
	return XBool(result)
}
