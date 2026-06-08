package goxt

type Equalable[T any] interface {
	Equal(other T) XBool
}

type Comparable[T any] interface {
	comparable
	Equalable[T]
}

type NumbersConstraints[T any] interface {
	Equalable[T]
	XInt | XInt8 | XInt16 | XInt32 | XInt64 | XByte | XUint | XUint8 | XUint16 | XUint32 | XUint64 | XFloat32 | XFloat64
}

type Integers interface {
	XInt | XInt8 | XInt16 | XInt32 | XInt64
}

type UIntegers interface {
	XByte | XUint | XUint8 | XUint16 | XUint32 | XUint64
}

type Floats interface {
	XFloat32 | XFloat64
}

type Complexes interface {
	XComplex64 | XComplex128
}

type Numbers interface {
	// XXX 临时去掉 Floats 的约束
	Integers | UIntegers /* | Floats */
}

type Appendable interface {
	Append(value XString)
	AppendWith(value XString, startIndex XInt, endIndex XInt)
	ToString() XString
}