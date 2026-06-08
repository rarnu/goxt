package goxt

import (
	"fmt"
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

func (l XList[T]) Map[R Equalable[R]](transform func(T) R) XList[R] {
	return l.MapTo(NewXList[R](), transform)
}

func (l XList[T]) MapTo[R Equalable[R]](destination XList[R], transform func(T) R) XList[R] {
	for _, item := range l {
		destination.Add(transform(item))
	}
	return destination
}

func (l XList[T]) MapIndexedTo[R Equalable[R]](destination XList[R], transform func(XInt, T) R) XList[R] {
	for index, item := range l {
		destination.Add(transform(XInt(index), item))
	}
	return destination
}

func (l XList[T]) MapIndexed[R Equalable[R]](transform func(XInt, T) R) XList[R] {
	return l.MapIndexedTo(NewXList[R](), transform)
}

func (l XList[T]) Distinct() XList[T] {
	return l.ToSet().ToList()
}

func (l XList[T]) DistinctBy[K Comparable[K]](selector func(T) K) XList[T] {
	st := NewXSet[K]()
	ret := NewXList[T]()
	for _, item := range l {
		key := selector(item)
		if st.Add(key) {
			ret.Add(item)
		}
	}
	return ret
}

func (l XList[T]) Intersect(other XList[T]) XSet[T] {
	st := NewXSet[T]()
	for _, item := range l {
		if other.Contains(item) {
			st.Add(item)
		}
	}
	return st
}

func (l XList[T]) Subtract(other XList[T]) XSet[T] {
	st := NewXSet[T]()
	for _, item := range l {
		if !other.Contains(item) {
			st.Add(item)
		}
	}
	return st
}

func (l XList[T]) Union(other XList[T]) XSet[T] {
	st := l.ToSet()
	st.AddAll(other)
	return st
}

func (l XList[T]) All(predicate func(T) XBool) XBool {
	if l.IsEmpty() {
		return true
	}
	for _, item := range l {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (l XList[T]) Any(predicate func(T) XBool) XBool {
	if l.IsEmpty() {
		return false
	}
	for _, item := range l {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (l XList[T]) Count(predicate func(T) XBool) XInt {
	count := 0
	for _, item := range l {
		if predicate(item) {
			count++
		}
	}
	return XInt(count)
}

func (l XList[T]) None(predicate func(T) XBool) XBool {
	if l.IsEmpty() {
		return true
	}
	for _, item := range l {
		if predicate(item) {
			return false
		}
	}
	return true
}

func (l XList[T]) Fold[R Equalable[R]](initial R, operation func(R, T) R) R {
	ret := initial
	for _, item := range l {
		ret = operation(ret, item)
	}
	return ret
}

func (l XList[T]) FoldIndexed[R Equalable[R]](initial R, operation func(XInt, R, T) R) R {
	ret := initial
	for index, item := range l {
		ret = operation(XInt(index), ret, item)
	}
	return ret
}

func (l XList[T]) MaxWith(comparator func(T, T) XInt) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	for _, item := range l[1:] {
		if comparator(item, m) > 0 {
			m = item
		}
	}
	return m
}

func (l XList[T]) MaxWithOrNil(comparator func(T, T) XInt) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	for _, item := range l[1:] {
		if comparator(item, m) > 0 {
			m = item
		}
	}
	return &m
}

func (l XList[T]) MinWith(comparator func(T, T) XInt) T {
	if l.IsEmpty() {
		panic("Collection is empty.")
	}
	m := l[0]
	for _, item := range l[1:] {
		if comparator(item, m) < 0 {
			m = item
		}
	}
	return m
}

func (l XList[T]) MinWithOrNil(comparator func(T, T) XInt) *T {
	if l.IsEmpty() {
		return nil
	}
	m := l[0]
	for _, item := range l[1:] {
		if comparator(item, m) < 0 {
			m = item
		}
	}
	return &m
}

func (l XList[T]) Reduce(operation func(T, T) T) T {
	if l.IsEmpty() {
		panic("Empty collection can't be reduced.")
	}
	acc := l[0]
	for _, item := range l[1:] {
		acc = operation(acc, item)
	}
	return acc
}

func (l XList[T]) ReduceOrNil(operation func(T, T) T) *T {
	if l.IsEmpty() {
		return nil
	}
	acc := l[0]
	for _, item := range l[1:] {
		acc = operation(acc, item)
	}
	return &acc
}

func (l XList[T]) ReduceIndexed(operation func(XInt, T, T) T) T {
	if l.IsEmpty() {
		panic("Empty collection can't be reduced.")
	}
	acc := l[0]
	for index, item := range l[1:] {
		acc = operation(XInt(index), acc, item)
	}
	return acc
}

func (l XList[T]) ReduceIndexedOrNil(operation func(XInt, T, T) T) *T {
	if l.IsEmpty() {
		return nil
	}
	acc := l[0]
	for index, item := range l[1:] {
		acc = operation(XInt(index), acc, item)
	}
	return &acc
}

func (l XList[T]) RunningFold[R Equalable[R]](initial R, operation func(R, T) R) XList[R] {
	if l.IsEmpty() {
		return NewXListWithElements(initial)
	}
	ret := make(XList[R], 0, l.Size()+1)
	ret.Add(initial)
	acc := initial
	for _, item := range l {
		acc = operation(acc, item)
		ret.Add(acc)
	}
	return ret
}

func (l XList[T]) RunningFoldIndexed[R Equalable[R]](initial R, operation func(XInt, R, T) R) XList[R] {
	if l.IsEmpty() {
		return NewXListWithElements(initial)
	}
	ret := make(XList[R], 0, l.Size()+1)
	ret.Add(initial)
	acc := initial
	for index, item := range l {
		acc = operation(XInt(index), acc, item)
		ret.Add(acc)
	}
	return ret
}

func (l XList[T]) RunningReduce(operation func(T, T) T) XList[T] {
	if l.IsEmpty() {
		return EmptyXList[T]()
	}

	acc := l[0]
	ret := make(XList[T], 0, l.Size())
	ret.Add(acc)
	for _, item := range l[1:] {
		acc = operation(acc, item)
		ret.Add(acc)
	}
	return ret
}

func (l XList[T]) RunningReduceIndexed(operation func(XInt, T, T) T) XList[T] {
	if l.IsEmpty() {
		return EmptyXList[T]()
	}

	acc := l[0]
	ret := make(XList[T], 0, l.Size())
	ret.Add(acc)
	for index, item := range l[1:] {
		acc = operation(XInt(index+1), acc, item)
		ret.Add(acc)
	}
	return ret
}

func (l XList[T]) Partition(predicate func(T) XBool) XPair[XList[T], XList[T]] {
	first := NewXList[T]()
	second := NewXList[T]()
	for _, item := range l {
		if predicate(item) {
			first.Add(item)
		} else {
			second.Add(item)
		}
	}
	return XPair[XList[T], XList[T]]{First: first, Second: second}
}

func (l XList[T]) JoinTo[A Appendable](buffer A, separator XString, prefix XString, postfix XString, limit XInt, truncated XString, transform func(T) XString) A {
	buffer.Append(prefix)
	count := XInt(0)
	for _, item := range l {
		count++
		if count > 1 {
			buffer.Append(separator)
		}
		if limit < 0 || count <= limit {
			buffer.Append(transform(item))
		} else {
			break
		}
	}
	if limit >= 0 && count > limit {
		buffer.Append(truncated)
	}
	buffer.Append(postfix)
	return buffer
}

func (l XList[T]) JoinToString(separator XString, prefix XString, postfix XString, transform func(T) XString) XString {
	var a Appendable = new(XString(""))
	return l.JoinTo(a, separator, prefix, postfix, -1, "...", transform).ToString()
}

func (l XList[T]) JoinToStringWithAllDefault(separator XString) XString {
	var a Appendable = new(XString(""))
	return l.JoinTo(a, separator, "", "", -1, "...", func(item T) XString {
		return XString(fmt.Sprintf("%v", item))
	}).ToString()
}

func (l XList[T]) JoinToStringWithDefaultTransform(separator XString, prefix XString, postfix XString) XString {
	var a Appendable = new(XString(""))
	return l.JoinTo(a, separator, prefix, postfix, -1, "...", func(item T) XString {
		return XString(fmt.Sprintf("%v", item))
	}).ToString()
}

/*
func (l XList[T]) Zip[R Equalable[R]](other XList[R]) XList[XPair[T, R]]  {
	return nil
}
 */

/*
func (l XList[T]) Chunked(size XInt) XList[XList[T]] {
	// TODO Chunked
	return nil
}

public fun <T, R> Iterable<T>.chunked(size: Int, transform: (List<T>) -> R): List<R>

func (l XList[T]) Windowed(size XInt, step XInt, partialWindows XBool) XList[XList[T]]  {
	return nil
}

*/

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

/*
func (l XList[T]) GroupByTo[K Comparable[K]](destination XMap[K, XList[T]], selector func(T) K) XMap[K, XList[T]] {
	// TODO 待实现
	return nil
}

*/

// public fun <T> Iterable<Iterable<T>>.flatten(): List<T>

// public fun <T, R> Iterable<Pair<T, R>>.unzip(): Pair<List<T>, List<R>>
