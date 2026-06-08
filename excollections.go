package goxt

import (
	"math/rand"
	"time"
)

func (l XList[T]) Find(predicate func(T) XBool) *T {
	return l.FirstOrNilWithPredicate(predicate)
}

func (l XList[T]) FindLast(predicate func(T) XBool) *T {
	return l.LastOrNilWithPredicate(predicate)
}

func (l XList[T]) FirstWithPredicate(predicate func(T) XBool) T {
	for _, item := range l {
		if predicate(item) {
			return item
		}
	}
	panic("Collection contains no element matching the predicate.")
}

func (l XList[T]) GetOrNil(index XInt) *T {
	if index >= 0 && index < l.Size() {
		return &l[index]
	}
	return nil
}

func (l XList[T]) FirstOrNilWithPredicate(predicate func(T) XBool) *T {
	for _, item := range l {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

func (l XList[T]) IndexOfFirstWithPredicate(predicate func(T) XBool) XInt {
	for index, item := range l {
		if predicate(item) {
			return XInt(index)
		}
	}
	return -1
}

func (l XList[T]) IndexOfLastWithPredicate(predicate func(T) XBool) XInt {
	for index := l.Size() - 1; index >= 0; index-- {
		item := l[index]
		if predicate(item) {
			return index
		}
	}
	return -1
}

func (l XList[T]) LastWithPredicate(predicate func(T) XBool) T {
	for index := l.Size() - 1; index >= 0; index-- {
		item := l[index]
		if predicate(item) {
			return item
		}
	}
	panic("Collection contains no element matching the predicate.")
}

func (l XList[T]) LastOrNilWithPredicate(predicate func(T) XBool) *T {
	for index := l.Size() - 1; index >= 0; index-- {
		item := l[index]
		if predicate(item) {
			return &item
		}
	}
	return nil
}

func (l XList[T]) Random() T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return l[rand.Intn(len(l))]
}

func (l XList[T]) RandomOrNil() *T {
	if l.IsEmpty() {
		return nil
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &l[rand.Intn(len(l))]
}

func (l XList[T]) Drop(n XInt) XList[T] {
	if n >= l.Size() {
		return NewXList[T]()
	}
	return l.SubList(n, l.Size())
}
func (l XList[T]) DropLast(n XInt) XList[T] {
	if n >= l.Size() {
		return NewXList[T]()
	}
	return l.SubList(0, l.Size()-n)
}
func (l XList[T]) DropWithPredicate(predicate func(T) XBool) XList[T] {
	if l == nil {
		return nil
	}
	src := l
	result := make(XList[T], 0, len(src))
	for _, v := range src {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func (l XList[T]) DropLastWithPredicate(predicate func(T) XBool) XList[T] {
	if l == nil {
		return nil
	}
	src := l
	// 从后向前查找最后一个满足条件的索引
	lastIdx := -1
	for i := len(src) - 1; i >= 0; i-- {
		if predicate(src[i]) {
			lastIdx = i
			break
		}
	}
	if lastIdx == -1 {
		// 没有找到，返回原列表副本
		newList := make(XList[T], len(src))
		copy(newList, src)
		return newList
	}
	// 构建新切片，跳过 lastIdx
	newList := make(XList[T], 0, len(src)-1)
	newList = append(newList, src[:lastIdx]...)
	newList = append(newList, src[lastIdx+1:]...)
	return newList
}

func (l XList[T]) FilterTo(destination XList[T], predicate func(T) XBool) XList[T] {
	for _, item := range l {
		if predicate(item) {
			destination.Add(item)
		}
	}
	return destination
}

func (l XList[T]) Filter(predicate func(T) XBool) XList[T] {
	return l.FilterTo(NewXList[T](), predicate)
}

func (l XList[T]) FilterNotTo(destination XList[T], predicate func(T) XBool) XList[T] {
	for _, item := range l {
		if !predicate(item) {
			destination.Add(item)
		}
	}
	return destination
}

func (l XList[T]) FilterNot(predicate func(T) XBool) XList[T] {
	return l.FilterNotTo(NewXList[T](), predicate)
}

func (l XList[T]) FilterIndexedTo(destination XList[T], predicate func(index XInt, item T) XBool) XList[T] {
	for index, item := range l {
		if predicate(XInt(index), item) {
			destination.Add(item)
		}
	}
	return destination
}

func (l XList[T]) FilterIndexed(predicate func(index XInt, item T) XBool) XList[T] {
	return l.FilterIndexedTo(NewXList[T](), predicate)
}

func (l XList[T]) Take(n XInt) XList[T] {
	if n >= l.Size() {
		return l
	}
	return l.SubList(0, n)
}

func (l XList[T]) TakeLast(n XInt) XList[T] {
	if n >= l.Size() {
		return l
	}
	return l.SubList(l.Size()-n, l.Size())
}

func (l XList[T]) TakeWithPredicate(predicate func(T) XBool) XList[T] {
	if l == nil {
		return nil
	}
	src := l
	for i, v := range src {
		if !predicate(v) {
			// 返回前 i 个元素（复制）
			result := make(XList[T], i)
			copy(result, src[:i])
			return result
		}
	}
	// 所有元素都满足，返回整个列表的副本
	result := make(XList[T], len(src))
	copy(result, src)
	return result
}

func (l XList[T]) TakeLastWithPredicate(predicate func(T) XBool) XList[T] {
	if l == nil {
		return nil
	}
	src := l
	// 从后向前找第一个不满足 predicate 的位置
	startIdx := len(src)
	for i := len(src) - 1; i >= 0; i-- {
		if !predicate(src[i]) {
			break
		}
		startIdx = i
	}
	if startIdx == len(src) {
		// 没有任何元素满足（最后一个就不满足），返回空切片
		return XList[T]{}
	}
	// 切片 [startIdx:] 即为最长后缀
	result := make(XList[T], len(src)-startIdx)
	copy(result, src[startIdx:])
	return result
}

func (l XList[T]) AssociateTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], transform func(T) XMapEntry[K, V]) XMap[K, V] {
	for _, item := range l {
		pair := transform(item)
		destination[pair.Key] = pair.Value
	}
	return destination
}

func (l XList[T]) Associate[K Comparable[K], V Equalable[V]](transform func(T) XMapEntry[K, V]) XMap[K, V] {
	dest := make(XMap[K, V], l.Size())
	return l.AssociateTo(dest, transform)
}

func (l XList[T]) AssociateByWithValueTo[K Comparable[K], V Equalable[V]](destination XMap[K, V], selector func(T) K, transform func(T) V) XMap[K, V] {
	for _, item := range l {
		key := selector(item)
		value := transform(item)
		destination[key] = value
	}
	return destination
}

func (l XList[T]) AssociateByWithValue[K Comparable[K], V Equalable[V]](selector func(T) K, transform func(T) V) XMap[K, V] {
	dest := make(XMap[K, V], l.Size())
	return l.AssociateByWithValueTo(dest, selector, transform)
}

func (l XList[T]) ToSet() XSet[T] {
	ret := make(XSet[T], l.Size())
	for _, item := range l {
		ret.Add(item)
	}
	return ret
}

func (l XList[T]) FlatMapTo[R Equalable[R]](destination XList[R], transform func(T) XList[R]) XList[R] {
	for _, item := range l {
		destination.AddAll(transform(item))
	}
	return destination
}

func (l XList[T]) FlatMap[R Equalable[R]](transform func(T) XList[R]) XList[R] {
	dest := make(XList[R], l.Size())
	return l.FlatMapTo(dest, transform)
}

func (l XList[T]) FlatMapIndexedTo[R Equalable[R]](destination XList[R], transform func(XInt, T) XList[R]) XList[R] {
	for index, item := range l {
		destination.AddAll(transform(XInt(index), item))
	}
	return destination
}

func (l XList[T]) FlatMapIndexed[R Equalable[R]](transform func(XInt, T) XList[R]) XList[R] {
	dest := make(XList[R], l.Size())
	return l.FlatMapIndexedTo(dest, transform)
}

func (l XList[T]) GroupByTo[K comparable](destination map[K]XList[T], selector func(T) K) map[K]XList[T] {
	// TODO 待实现
	return nil
}

func (l XList[T]) Map[R Equalable[R]](transform func(T) R) XList[R] {
	return l.MapTo(NewXList[R](), transform)
}

func (l XList[T]) MapTo[R Equalable[R]](destination XList[R], transform func(T) R) XList[R] {
	for _, item := range l {
		destination.Add(transform(item))
	}
	return destination
}

func (s XSet[T]) ToList() XList[T] {
	ret := make(XList[T], s.Size())
	for _, item := range s {
		ret.Add(item)
	}
	return ret
}

func (m XMap[K, V]) ToList() XList[XMapEntry[K, V]] {
	ret := make(XList[XMapEntry[K, V]], m.Size())
	for _, entry := range m.Entries() {
		ret.Add(entry)
	}
	return ret
}

/*
func (l XList[T]) AssociateByTo[K Comparable[K]](destination XMap[K, T], selector func(T) K) XMap[K, T] {
	for _, item := range l {
		key := selector(item)
		destination[key] = item
	}
	return destination
}
 */


/*
func (l XList[T]) AssociateBy[K comparable](selector func(T) K) XMap[K, T] {
	dest := make(XMap[K, T], l.Size())
	return l.AssociateByTo(dest, selector)
}

*/

/*
func (l XList[T]) AssociateWithTo[V comparable](destination XMap[T, V], selector func(T) V) XMap[T, V] {
	for _, item := range l {
		value := selector(item)
		destination[item] = value
	}
	return destination
}

*/

/*
func (l XList[T]) AssociateWith[V comparable](selector func(T) V) XMap[T, V] {
	dest := make(XMap[T, V], l.Size())
	return l.AssociateWithTo(dest, selector)
}

*/

// public fun <T> Iterable<Iterable<T>>.flatten(): List<T>

// public fun <T, R> Iterable<Pair<T, R>>.unzip(): Pair<List<T>, List<R>>
