package goxt

import (
	"fmt"
	"strconv"
)

type XInt int
type XInt8 int8
type XInt16 int16
type XInt32 int32
type XInt64 int64
type XUint uint
type XUint8 uint8
type XUint16 uint16
type XUint32 uint32
type XUint64 uint64
type XByte byte
type XFloat32 float32
type XFloat64 float64
type XComplex64 complex64
type XComplex128 complex128

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

type Number[T Numbers] interface {
	ToFloat64() XFloat64
	ToFloat32() XFloat32
	ToInt64() XInt64
	ToInt32() XInt32
	ToInt() XInt
	ToInt16() XInt16
	ToInt8() XInt8
	ToRune() XRune
	ToByte() XByte
	ToUint64() XUint64
	ToUint32() XUint32
	ToUint() XUint
	ToUint16() XUint16
	ToUint8() XUint8
	ToString() XString
}

var (
	_ Number[XInt]    = (*XInt)(nil)
	_ Number[XInt8]   = (*XInt8)(nil)
	_ Number[XInt16]  = (*XInt16)(nil)
	_ Number[XInt32]  = (*XInt32)(nil)
	_ Number[XInt64]  = (*XInt64)(nil)
	_ Number[XUint]   = (*XUint)(nil)
	_ Number[XUint8]  = (*XUint8)(nil)
	_ Number[XUint16] = (*XUint16)(nil)
	_ Number[XUint32] = (*XUint32)(nil)
	_ Number[XUint64] = (*XUint64)(nil)
	_ Number[XByte]   = (*XByte)(nil)
	// XXX 临时去掉 Floats 的约束
	// _ Float[XFloat32] = (*XFloat32)(nil)
	// _ Float[XFloat64] = (*XFloat64)(nil)
)

func (i XInt) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XInt) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XInt) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XInt) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XInt) ToInt() XInt {
	return i
}
func (i XInt) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XInt) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XInt) ToRune() XRune {
	return XRune(i)
}
func (i XInt) ToByte() XByte {
	return XByte(i)
}
func (i XInt) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XInt) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XInt) ToUint() XUint {
	return XUint(i)
}
func (i XInt) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XInt) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XInt) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XInt8) ToFloat64() XFloat64 {
	return XFloat64(i)
}

func (i XInt8) ToFloat32() XFloat32 {
	return XFloat32(i)
}

func (i XInt8) ToInt64() XInt64 {
	return XInt64(i)
}

func (i XInt8) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XInt8) ToInt() XInt {
	return XInt(i)
}
func (i XInt8) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XInt8) ToInt8() XInt8 {
	return i
}
func (i XInt8) ToRune() XRune {
	return XRune(i)
}
func (i XInt8) ToByte() XByte {
	return XByte(i)
}
func (i XInt8) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XInt8) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XInt8) ToUint() XUint {
	return XUint(i)
}
func (i XInt8) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XInt8) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XInt8) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XInt16) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XInt16) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XInt16) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XInt16) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XInt16) ToInt() XInt {
	return XInt(i)
}
func (i XInt16) ToInt16() XInt16 {
	return i
}
func (i XInt16) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XInt16) ToRune() XRune {
	return XRune(i)
}
func (i XInt16) ToByte() XByte {
	return XByte(i)
}
func (i XInt16) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XInt16) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XInt16) ToUint() XUint {
	return XUint(i)
}
func (i XInt16) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XInt16) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XInt16) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XInt32) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XInt32) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XInt32) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XInt32) ToInt32() XInt32 {
	return i
}
func (i XInt32) ToInt() XInt {
	return XInt(i)
}
func (i XInt32) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XInt32) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XInt32) ToRune() XRune {
	return XRune(i)
}
func (i XInt32) ToByte() XByte {
	return XByte(i)
}
func (i XInt32) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XInt32) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XInt32) ToUint() XUint {
	return XUint(i)
}
func (i XInt32) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XInt32) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XInt32) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XInt64) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XInt64) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XInt64) ToInt64() XInt64 {
	return i
}
func (i XInt64) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XInt64) ToInt() XInt {
	return XInt(i)
}
func (i XInt64) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XInt64) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XInt64) ToRune() XRune {
	return XRune(i)
}
func (i XInt64) ToByte() XByte {
	return XByte(i)
}
func (i XInt64) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XInt64) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XInt64) ToUint() XUint {
	return XUint(i)
}
func (i XInt64) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XInt64) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XInt64) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XUint) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XUint) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XUint) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XUint) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XUint) ToInt() XInt {
	return XInt(i)
}
func (i XUint) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XUint) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XUint) ToRune() XRune {
	return XRune(i)
}
func (i XUint) ToByte() XByte {
	return XByte(i)
}
func (i XUint) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XUint) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XUint) ToUint() XUint {
	return i
}
func (i XUint) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XUint) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XUint) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

// ==================

func (i XUint8) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XUint8) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XUint8) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XUint8) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XUint8) ToInt() XInt {
	return XInt(i)
}
func (i XUint8) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XUint8) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XUint8) ToRune() XRune {
	return XRune(i)
}
func (i XUint8) ToByte() XByte {
	return XByte(i)
}
func (i XUint8) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XUint8) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XUint8) ToUint() XUint {
	return XUint(i)
}
func (i XUint8) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XUint8) ToUint8() XUint8 {
	return i
}
func (i XUint8) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XUint16) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XUint16) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XUint16) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XUint16) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XUint16) ToInt() XInt {
	return XInt(i)
}
func (i XUint16) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XUint16) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XUint16) ToRune() XRune {
	return XRune(i)
}
func (i XUint16) ToByte() XByte {
	return XByte(i)
}
func (i XUint16) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XUint16) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XUint16) ToUint() XUint {
	return XUint(i)
}
func (i XUint16) ToUint16() XUint16 {
	return i
}
func (i XUint16) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XUint16) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XUint32) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XUint32) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XUint32) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XUint32) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XUint32) ToInt() XInt {
	return XInt(i)
}
func (i XUint32) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XUint32) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XUint32) ToRune() XRune {
	return XRune(i)
}
func (i XUint32) ToByte() XByte {
	return XByte(i)
}
func (i XUint32) ToUint64() XUint64 {
	return XUint64(i)
}
func (i XUint32) ToUint32() XUint32 {
	return i
}
func (i XUint32) ToUint() XUint {
	return XUint(i)
}
func (i XUint32) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XUint32) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XUint32) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (i XUint64) ToFloat64() XFloat64 {
	return XFloat64(i)
}
func (i XUint64) ToFloat32() XFloat32 {
	return XFloat32(i)
}
func (i XUint64) ToInt64() XInt64 {
	return XInt64(i)
}
func (i XUint64) ToInt32() XInt32 {
	return XInt32(i)
}
func (i XUint64) ToInt() XInt {
	return XInt(i)
}
func (i XUint64) ToInt16() XInt16 {
	return XInt16(i)
}
func (i XUint64) ToInt8() XInt8 {
	return XInt8(i)
}
func (i XUint64) ToRune() XRune {
	return XRune(i)
}
func (i XUint64) ToByte() XByte {
	return XByte(i)
}
func (i XUint64) ToUint64() XUint64 {
	return i
}
func (i XUint64) ToUint32() XUint32 {
	return XUint32(i)
}
func (i XUint64) ToUint() XUint {
	return XUint(i)
}
func (i XUint64) ToUint16() XUint16 {
	return XUint16(i)
}
func (i XUint64) ToUint8() XUint8 {
	return XUint8(i)
}
func (i XUint64) ToString() XString {
	return XString(fmt.Sprintf("%d", i))
}

func (b XByte) ToFloat64() XFloat64 {
	return XFloat64(b)
}
func (b XByte) ToFloat32() XFloat32 {
	return XFloat32(b)
}
func (b XByte) ToInt64() XInt64 {
	return XInt64(b)
}
func (b XByte) ToInt32() XInt32 {
	return XInt32(b)
}
func (b XByte) ToInt() XInt {
	return XInt(b)
}
func (b XByte) ToInt16() XInt16 {
	return XInt16(b)
}
func (b XByte) ToInt8() XInt8 {
	return XInt8(b)
}
func (b XByte) ToRune() XRune {
	return XRune(b)
}
func (b XByte) ToByte() XByte {
	return b
}
func (b XByte) ToUint64() XUint64 {
	return XUint64(b)
}
func (b XByte) ToUint32() XUint32 {
	return XUint32(b)
}
func (b XByte) ToUint() XUint {
	return XUint(b)
}
func (b XByte) ToUint16() XUint16 {
	return XUint16(b)
}
func (b XByte) ToUint8() XUint8 {
	return XUint8(b)
}
func (b XByte) ToString() XString {
	return XString(fmt.Sprintf("%d", b))
}

func (f XFloat32) ToFloat64() XFloat64 {
	return XFloat64(f)
}
func (f XFloat32) ToFloat32() XFloat32 {
	return f
}
func (f XFloat32) ToInt64() XInt64 {
	return XInt64(f)
}
func (f XFloat32) ToInt32() XInt32 {
	return XInt32(f)
}
func (f XFloat32) ToInt() XInt {
	return XInt(f)
}
func (f XFloat32) ToInt16() XInt16 {
	return XInt16(f)
}
func (f XFloat32) ToInt8() XInt8 {
	return XInt8(f)
}
func (f XFloat32) ToRune() XRune {
	return XRune(f)
}
func (f XFloat32) ToByte() XByte {
	return XByte(f)
}
func (f XFloat32) ToUint64() XUint64 {
	return XUint64(f)
}
func (f XFloat32) ToUint32() XUint32 {
	return XUint32(f)
}
func (f XFloat32) ToUint() XUint {
	return XUint(f)
}
func (f XFloat32) ToUint16() XUint16 {
	return XUint16(f)
}
func (f XFloat32) ToUint8() XUint8 {
	return XUint8(f)
}
func (f XFloat32) ToString() XString {
	return XString(strconv.FormatFloat(float64(f), 'f', -1, 32))
}

func (f XFloat64) ToFloat64() XFloat64 {
	return f
}
func (f XFloat64) ToFloat32() XFloat32 {
	return XFloat32(f)
}
func (f XFloat64) ToInt64() XInt64 {
	return XInt64(f)
}
func (f XFloat64) ToInt32() XInt32 {
	return XInt32(f)
}
func (f XFloat64) ToInt() XInt {
	return XInt(f)
}
func (f XFloat64) ToInt16() XInt16 {
	return XInt16(f)
}
func (f XFloat64) ToInt8() XInt8 {
	return XInt8(f)
}
func (f XFloat64) ToRune() XRune {
	return XRune(f)
}
func (f XFloat64) ToByte() XByte {
	return XByte(f)
}
func (f XFloat64) ToUint64() XUint64 {
	return XUint64(f)
}
func (f XFloat64) ToUint32() XUint32 {
	return XUint32(f)
}
func (f XFloat64) ToUint() XUint {
	return XUint(f)
}
func (f XFloat64) ToUint16() XUint16 {
	return XUint16(f)
}
func (f XFloat64) ToUint8() XUint8 {
	return XUint8(f)
}
func (f XFloat64) ToString() XString {
	return XString(strconv.FormatFloat(float64(f), 'f', -1, 64))
}
