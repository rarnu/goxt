package goxt

type XBool bool

func (b XBool) ToString() XString {
	if b {
		return "true"
	}
	return "false"
}
