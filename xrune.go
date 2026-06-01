package goxt

type XRune rune

const (
	XRuneMinValue = '\u0000'
	XRuneMaxValue = '\uFFFF'
)

func (r XRune) ToByte() XByte {
	return XByte(r)
}

func (r XRune) ToInt() XInt {
	return XInt(r)
}

func (r XRune) ToInt8() XInt8 {
	return XInt8(r)
}

func (r XRune) ToInt16() XInt16 {
	return XInt16(r)
}

func (r XRune) ToInt32() XInt32 {
	return XInt32(r)
}

func (r XRune) ToInt64() XInt64 {
	return XInt64(r)
}

func (r XRune) ToUint() XUint {
	return XUint(r)
}

func (r XRune) ToUint8() XUint8 {
	return XUint8(r)
}

func (r XRune) ToUint16() XUint16 {
	return XUint16(r)
}

func (r XRune) ToUint32() XUint32 {
	return XUint32(r)
}

func (r XRune) ToUint64() XUint64 {
	return XUint64(r)
}

func (r XRune) ToFloat32() XFloat32 {
	return XFloat32(r)
}

func (r XRune) ToFloat64() XFloat64 {
	return XFloat64(r)
}

// Char Creates a Char with the specified [code], or throws an exception if the [code] is out of `Char.MIN_VALUE.code to Char.MAX_VALUE.code`.
func Char(code XInt) XRune {
	if code < XRuneMinValue || code > XRuneMaxValue {
		return XRuneMinValue
	}
	return XRune(code)
}

func (r XRune) Code() XInt {
	return XInt(r)
}

// public expect fun CharArray.concatToString(): String

// public expect fun CharArray.concatToString(startIndex: Int = 0, endIndex: Int = this.size): String
