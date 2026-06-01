package goxt

func IfThen[T any](condition XBool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
