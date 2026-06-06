package goxt

import (
	"math"
	"strconv"
)

func (i XInt) RangeToOpen(to XInt) XList[XInt] {
	ret := make([]XInt, to-i)
	for j := i; j < to; j++ {
		ret[j-i] = j
	}
	return ret
}

func (i XInt) RangeToClose(to XInt) XList[XInt] {
	ret := make([]XInt, to-i+1)
	for j := i; j <= to; j++ {
		ret[j-i] = j
	}
	return ret
}

func (i XInt) ToStringRadix(radix int) XString {
	return XString(strconv.FormatInt(int64(i), radix))
}

func (i XInt) Absolute() XInt {
	return XInt(math.Abs(float64(i)))
}

func (i XInt8) ToStringRadix(radix int) XString {
	return XString(strconv.FormatInt(int64(i), radix))
}

func (i XInt8) Absolute() XInt8 {
	return XInt8(math.Abs(float64(i)))
}

func (i XInt16) ToStringRadix(radix int) XString {
	return XString(strconv.FormatInt(int64(i), radix))
}

func (i XInt16) Absolute() XInt16 {
	return XInt16(math.Abs(float64(i)))
}

func (i XInt32) ToStringRadix(radix int) XString {
	return XString(strconv.FormatInt(int64(i), radix))
}

func (i XInt32) Absolute() XInt32 {
	return XInt32(math.Abs(float64(i)))
}

func (i XInt64) ToStringRadix(radix int) XString {
	return XString(strconv.FormatInt(int64(i), radix))
}

func (i XInt64) Absolute() XInt64 {
	return XInt64(math.Abs(float64(i)))
}

func (i XUint) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(i)), radix))
}

func (i XUint8) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(i)), radix))
}

func (i XUint16) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(i)), radix))
}

func (i XUint32) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(i)), radix))
}

func (i XUint64) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(i)), radix))
}

func (b XByte) ToStringRadix(radix int) XString {
	return XString(strconv.FormatUint(uint64(int64(b)), radix))
}
