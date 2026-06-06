package goxt

import "math"

func (f XFloat32) IsNaN() XBool {
	return XBool(math.IsNaN(float64(f)))
}

func (f XFloat32) IsInfinite() XBool {
	return XBool(math.IsInf(float64(f), 0))
}

func (f XFloat64) IsNaN() XBool {
	return XBool(math.IsNaN(float64(f)))
}

func (f XFloat64) IsInfinite() XBool {
	return XBool(math.IsInf(float64(f), 0))
}

func (f XFloat32) Pow[T Numbers](x T) XFloat64  {
	return XFloat64(math.Pow(float64(f), float64(x)))
}

func (f XFloat64) Pow[T Numbers](x T) XFloat64  {
	return XFloat64(math.Pow(float64(f), float64(x)))
}

func (f XFloat32) Absolute() XFloat32 {
	return XFloat32(math.Abs(float64(f)))
}

func (f XFloat64) Absolute() XFloat64 {
	return XFloat64(math.Abs(float64(f)))
}

