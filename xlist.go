package goxt

type XList[T any] []T

func (l XList[T]) Size() XInt {
	return XInt(len(l))
}

func (l XList[T]) Map[R any](transform func(T) R) XList[R]  {
	ret := make(XList[R], len(l))
	for index, item := range l {
		ret[index] = transform(item)
	}
	return ret
}