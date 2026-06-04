package goxt

func (l *XList[T]) Map[R comparable](transform func(T) R) XList[R] {
	return l.MapTo(NewXList[R](), transform)
}

func (l *XList[T]) MapTo[R comparable](destination XList[R], transform func(T) R) XList[R] {
	for _, item := range *l {
		destination.Add(transform(item))
	}
	return destination
}
