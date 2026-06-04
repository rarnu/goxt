package goxt

import (
	"math"
	"strconv"
)

//type XInt8 int8
//type XInt16 int16
//type XInt32 int32
//type XInt64 int64
//type XUint uint
//type XUint8 uint8
//type XUint16 uint16
//type XUint32 uint32
//type XUint64 uint64
//type XByte byte

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

func (i XInt) ToStringRadix(radix int) string {
	return strconv.FormatInt(int64(i), radix)
}

func (i XInt) Absolute() XInt {
	return XInt(math.Abs(float64(i)))
}

func (i XInt8) ToStringRadix(radix int) string {
	return strconv.FormatInt(int64(i), radix)
}

func (i XInt8) Absolute() XInt8 {
	return XInt8(math.Abs(float64(i)))
}

func (i XInt16) ToStringRadix(radix int) string {
	return strconv.FormatInt(int64(i), radix)
}

func (i XInt16) Absolute() XInt16 {
	return XInt16(math.Abs(float64(i)))
}

func (i XInt32) ToStringRadix(radix int) string {
	return strconv.FormatInt(int64(i), radix)
}

func (i XInt32) Absolute() XInt32 {
	return XInt32(math.Abs(float64(i)))
}

func (i XInt64) ToStringRadix(radix int) string {
	return strconv.FormatInt(int64(i), radix)
}

func (i XInt64) Absolute() XInt64 {
	return XInt64(math.Abs(float64(i)))
}

func (i XUint) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

func (i XUint8) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

func (i XUint16) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

func (i XUint32) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

func (i XUint64) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

func (i XByte) ToStringRadix(radix int) string {
	return strconv.FormatUint(uint64(int64(i)), radix)
}

