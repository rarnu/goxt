package goxt

import (
	"math/rand"
	"time"
)

type XList[T comparable] []T

func NewXList[T comparable]() XList[T] {
	return XList[T]{}
}

func NewXListWithSize[T comparable](size XInt) XList[T] {
	return make(XList[T], size)
}

func NewXListWithElements[T comparable](elements ...T) XList[T] {
	return elements
}

func NewXListWithInit[T comparable](size XInt, init func(XInt) T) XList[T] {
	ret := make(XList[T], size)
	for i := 0; i < int(size); i++ {
		ret[i] = init(XInt(i))
	}
	return ret
}

func EmptyXList[T comparable]() XList[T] {
	return make(XList[T], 0)
}

func (l *XList[T]) Size() XInt {
	return XInt(len(*l))
}

func (l *XList[T]) IsEmpty() XBool {
	return l.Size() == 0
}

func (l *XList[T]) IsNotEmpty() XBool {
	return l.Size() != 0
}

func (l *XList[T]) Contains(element T) XBool {
	for _, item := range *l {
		if item == element {
			return true
		}
	}
	return false
}

func (l *XList[T]) ContainsAll(elements XList[T]) XBool {
	for _, item := range elements {
		if !l.Contains(item) {
			return false
		}
	}
	return true
}

func (l *XList[T]) Add(element T) XBool {
	*l = append(*l, element)
	return true
}

func (l *XList[T]) Remove(element T) XBool {
	for i, item := range *l {
		if item == element {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return true
		}
	}
	return false
}

func (l *XList[T]) AddAll(elements XList[T]) XBool {
	*l = append(*l, elements...)
	return true
}

func (l *XList[T]) RemoveAll(elements XList[T]) XBool {
	removeMap := make(map[T]bool)
	for _, v := range elements {
		removeMap[v] = true
	}
	result := make([]T, 0, len(*l))
	for _, v := range *l {
		if !removeMap[v] {
			result = append(result, v)
		}
	}
	*l = result
	return true
}

func (l *XList[T]) RetainAll(elements XList[T]) XBool {
	retainMap := make(map[T]bool)
	for _, v := range elements {
		retainMap[v] = true
	}
	n := 0
	for _, v := range *l {
		if retainMap[v] {
			(*l)[n] = v
			n++
		}
	}
	result := (*l)[:n]
	*l = result
	return true
}

func (l *XList[T]) Clear() {
	*l = EmptyXList[T]()
}

func (l *XList[T]) IndexOf(element T) XInt {
	for i, item := range *l {
		if item == element {
			return XInt(i)
		}
	}
	return -1
}
func (l *XList[T]) LastIndexOf(element T) XInt {
	for i := l.Size() - 1; i >= 0; i-- {
		if (*l)[i] == element {
			return i
		}
	}
	return -1
}

func (l *XList[T]) SubList(fromIndex XInt, toIndex XInt) XList[T] {
	return (*l)[fromIndex:toIndex]
}

func (l *XList[T]) Insert(index XInt, element T) XBool {
	list := *l
	if index < 0 || index > list.Size() {
		return false
	}
	newList := make([]T, 0, len(list)+1)
	newList = append(newList, list[:index]...)
	newList = append(newList, element)
	newList = append(newList, list[index:]...)
	*l = newList
	return true
}
func (l *XList[T]) RemoveAt(index XInt) *T {
	list := *l
	if index < 0 || index >= list.Size() {
		return nil
	}
	removed := list[index]
	*l = append(list[:index], list[index+1:]...)
	return new(removed)
}

func (l *XList[T]) First() T {
	if l.IsEmpty() {
		panic("List is empty.")
	}
	return (*l)[0]
}

func (l *XList[T]) FirstOrNull() *T {
	if l.IsEmpty() {
		return nil
	}
	return &(*l)[0]
}

func (l *XList[T]) Last() T {
	if l.IsEmpty() {
		panic("List is empty.")
	}
	return (*l)[l.Size()-1]
}

func (l *XList[T]) LastOrNull() *T {
	if l.IsEmpty() {
		return nil
	}
	return &(*l)[l.Size()-1]
}

func (l *XList[T]) AddFirst(element T) {
	l.Insert(0, element)
}

func (l *XList[T]) AddLast(element T) {
	l.Insert(l.Size(), element)
}

func (l *XList[T]) RemoveFirst() T {
	return *l.RemoveAt(0)
}

func (l *XList[T]) RemoveFirstOrNull() *T {
	return l.RemoveAt(0)
}

func (l *XList[T]) RemoveLast() T {
	return *l.RemoveAt(l.Size() - 1)
}

func (l *XList[T]) RemoveLastOrNull() *T {
	return l.RemoveAt(l.Size() - 1)
}

func (l *XList[T]) InsertAll(index XInt, elements XList[T]) XBool {
	list := *l
	if index < 0 || index > list.Size() {
		return false
	}
	if len(elements) == 0 {
		return true
	}
	newLen := len(list) + len(elements)
	if cap(list) >= newLen {
		*l = list[:newLen]
		newList := *l
		copy(newList[index+elements.Size():], newList[index:])
		// 复制要插入的元素
		copy(newList[index:], elements)
	} else {
		// 容量不足，创建新切片
		newList := make([]T, newLen)
		copy(newList, list[:index])
		copy(newList[index:], elements)
		copy(newList[index+elements.Size():], list[index:])
		*l = newList
	}
	return true
}

func (l *XList[T]) RemoveRange(fromIndex XInt, toIndex XInt) {
	list := *l
	// 检查索引有效性
	if fromIndex < 0 || toIndex > list.Size() || fromIndex >= toIndex {
		return
	}
	// 删除范围内的元素
	*l = append(list[:fromIndex], list[toIndex:]...)
}

func (l *XList[T]) LastIndex() XInt {
	return l.Size() - 1
}

func (l *XList[T]) ForEach(operation func(T)) {
	for _, item := range *l {
		operation(item)
	}
}

func (l *XList[T]) IfEmpty(defaultValue func() XList[T]) XList[T] {
	if l.IsEmpty() {
		return defaultValue()
	}
	return *l
}

func (l *XList[T]) RemoveAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make([]T, 0, len(*l))
	for _, elem := range *l {
		if !predicate(elem) {
			newItems = append(newItems, elem)
			result = true
		}
	}
	*l = newItems
	return XBool(result)
}
func (l *XList[T]) RetainAllWithPredicate(predicate func(element T) XBool) XBool {
	result := false
	newItems := make([]T, 0, len(*l))
	for _, elem := range *l {
		if predicate(elem) {
			newItems = append(newItems, elem)
			result = true
		}
	}
	*l = newItems
	return XBool(result)
}

func (l *XList[T]) Reversed() XList[T] {
	n := len(*l)
	reversed := make([]T, n)
	for i, elem := range *l {
		reversed[n-1-i] = elem
	}
	return reversed
}

func (l *XList[T]) Shuffled() XList[T] {
	n := len(*l)
	if n <= 1 {
		// 长度 0 或 1 时无需打乱，直接返回副本
		shuffled := make([]T, n)
		copy(shuffled, *l)
		return shuffled
	}

	// 创建新切片并复制原数据
	shuffled := make([]T, n)
	copy(shuffled, *l)

	// 使用当前时间作为随机种子
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
