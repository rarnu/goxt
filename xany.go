package goxt

type Equalable[T any] interface {
	Equal(other T) XBool
}

type Comparable[T any] interface {
	comparable
	Equalable[T]
}

type Nothing struct{}

func (n Nothing) Equal(_ Nothing) XBool {
	return true
}

type XAny[T any] struct {
	Value T
}

func NewXAny[T any](value T) XAny[T] {
	return XAny[T]{value}
}

func (a XAny[T]) Let[R any](block func(T) R) R {
	return block(a.Value)
}

func (a XAny[T]) Apply(block func(T)) XAny[T] {
	block(a.Value)
	return a
}

func (a XAny[T]) TakeIf(predicate func(T) bool) *XAny[T] {
	if predicate(a.Value) {
		return &a
	}
	return nil
}

func (a XAny[T]) TakeUnless(predicate func(T) bool) *XAny[T] {
	if !predicate(a.Value) {
		return &a
	}
	return nil
}

func (a XAny[T]) Repeat(times XInt, action func(XInt)) {
	for index := range times {
		action(index)
	}
}
