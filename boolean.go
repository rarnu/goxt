package goxt

type XBool bool

func (b XBool) Equal(other XBool) XBool {
	return b == other
}

func (b XBool) ToString() XString {
	if b {
		return "true"
	}
	return "false"
}
